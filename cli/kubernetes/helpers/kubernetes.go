/* SPDX-License-Identifier: AGPL-3.0-only
 * Copyright (c) Benedict Schlueter
 */

package helpers

import (
	"context"
	"time"

	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	metaAPI "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Client is the struct used to access kubernetes helpers.
type Client struct {
	client     kubernetes.Interface
	logger     *zap.Logger
	restClient *rest.Config
}

// NewClient returns a new kuberenetes client-go wrapper.
func NewClient(kubeconfigPath string, logger *zap.Logger) (*Client, error) {
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		return nil, err
	}
	// create the clientset
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return &Client{
		client:     client,
		logger:     logger,
		restClient: config,
	}, nil
}

// GetClient returns the kubernetes client.
func (k *Client) GetClient() kubernetes.Interface {
	return k.client
}

// CreateStatefulSetForUser creates want waits for the statefulSet.
func (k *Client) CreateStatefulSetForUser(ctx context.Context, challengeNamespace, userID string) error {
	exists, err := k.NamespaceExists(ctx, challengeNamespace)
	if err != nil {
		return err
	}
	if !exists {
		if err := k.CreateNamespace(ctx, challengeNamespace); err != nil {
			return err
		}
	}
	if err := k.CreateChallengeStatefulSet(ctx, challengeNamespace, userID); err != nil {
		return err
	}
	if err := k.WaitForStatefulSet(ctx, challengeNamespace, userID, 20*time.Second); err != nil {
		return err
	}
	return k.WaitForPodRunning(ctx, challengeNamespace, userID, 4*time.Minute)
}

// CreateSecret creates a secret.
func (k *Client) CreateSecret(ctx context.Context) error {
	secret := v1.Secret{
		TypeMeta: metaAPI.TypeMeta{
			Kind:       "Secret",
			APIVersion: v1.SchemeGroupVersion.Version,
		},
		ObjectMeta: metaAPI.ObjectMeta{
			Name:      "azure-disk-secret",
			Namespace: "kube-system",
		},
		Data: map[string][]byte{
			"azurestorageaccountname": []byte("delegatiokube"),
			"azurestorageaccountkey":  []byte("dDT6N/8laMdsF8lUsbOc6sKbFTZYWNCKt5YF6b4hqGMI+fOw4X7HsPJTTcUBYwlOjIi6viAQ38iB+AStzSxAlQ=="),
		},
	}

	_, err := k.client.CoreV1().Secrets("kube-system").Create(ctx, &secret, metaAPI.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}
