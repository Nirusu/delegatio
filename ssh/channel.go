/* SPDX-License-Identifier: AGPL-3.0-only
 * Copyright (c) Benedict Schlueter
 */

package main

import (
	"sync"

	"golang.org/x/crypto/ssh"
)

// SSHChannelServer is a wrapper around an ssh.Channel and ssh.Requests.
type SSHChannelServer struct {
	channel  ssh.Channel
	requests *ssh.Request
	ptyReq   *PtyRequestPayload
	wg       *sync.WaitGroup
	done     chan struct{}
}

func NewChannel(channel ssh.Channel, requests *ssh.Request) *SSHChannelServer {
	return &SSHChannelServer{
		channel:  channel,
		requests: requests,
	}
}
