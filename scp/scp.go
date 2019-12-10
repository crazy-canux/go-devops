package scp

import (
	"golang.org/x/crypto/ssh"
	"time"
)

// It has a default timeout of one minute.
func NewClient(host string, config *ssh.ClientConfig) Client {
	return NewConfigurer(host, config).Create()
}

// Returns a new scp.Client with provides host, ssh.ClientConfig and timeout
func NewClientWithTimeout(host string, config *ssh.ClientConfig, timeout time.Duration) Client {
	return NewConfigurer(host, config).Timeout(timeout).Create()
}