/* SPDX-License-Identifier: AGPL-3.0-only
 * Copyright (c) Benedict Schlueter
 */

package helpers

import (
	"bytes"
	"context"
	"errors"
	"net/http"

	"go.uber.org/zap"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"
)

// CreatePodPortForward creates a port forward to a pod.
func (k *Client) CreatePodPortForward(ctx context.Context, namespace, podName, podPort string) (chan<- struct{}, error) {
	roundTripper, upgrader, err := spdy.RoundTripperFor(k.restClient)
	if err != nil {
		k.logger.Error("failed to create round tripper", zap.Error(err))
		return nil, err
	}

	req := k.client.CoreV1().RESTClient().Post().Resource("pods").Name(podName).Namespace(namespace).SubResource("portforward")

	/* 	hostIP := strings.TrimLeft(k.restClient.Host, "htps:/")
	   	serverURL := url.URL{Scheme: "https", Path: path, Host: hostIP} */

	dialer := spdy.NewDialer(upgrader, &http.Client{Transport: roundTripper}, http.MethodPost, req.URL())
	stopChan, readyChan := make(chan struct{}, 1), make(chan struct{}, 1)
	out := new(bytes.Buffer)

	forwarder, err := portforward.New(dialer, []string{podPort}, stopChan, readyChan, out, out)
	if err != nil {
		k.logger.Error("failed to create port forwarder", zap.Error(err))
		return nil, err
	}
	readyChan <- struct{}{}
	ports, err := forwarder.GetPorts()
	if err != nil {
		k.logger.Error("failed to get ports", zap.Error(err))
		return nil, err
	}
	if len(ports) != 1 {
		k.logger.Info("failed to forward ports", zap.Int("ports", len(ports)))
		return nil, errors.New("more than one port to forward")
	}

	go func() {
		if err := forwarder.ForwardPorts(); err != nil {
			k.logger.Error("failed to forward ports", zap.Error(err))
		}
		k.logger.Info("port forwarding stopped", zap.String("pod", podName), zap.String("port", podPort))
	}()

	return stopChan, nil
}
