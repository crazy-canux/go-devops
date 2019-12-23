package ssh

import (
	"golang.org/x/crypto/ssh"
	"time"
)

// use username and password.
func NewClientWithBasicAuth(host, username, password string) Client {
	config, _ := PasswordKey(username, password)
	return NewConfigurerWithBasicAuth(host, password, &config).Create()
}

// user username and password with timeout
func NewClientWithBasicAuthAndTimeout(host, username, password string, timeout time.Duration) Client {
	config, _ := PasswordKey(username, password)
	return NewConfigurerWithBasicAuth(host, password, &config).Timeout(timeout).Create()
}

// It has a default timeout of one minute.
func NewClient(host string, config *ssh.ClientConfig) Client {
	return NewConfigurer(host, config).Create()
}

// Returns a new ssh.Client with provides host, ssh.ClientConfig and timeout
func NewClientWithTimeout(host string, config *ssh.ClientConfig, timeout time.Duration) Client {
	return NewConfigurer(host, config).Timeout(timeout).Create()
}