package ssh

import (
	"io/ioutil"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

//PrivateKey Loads a private and public key from "path" and returns a SSH ClientConfig to authenticate with the server
func PrivateKey(username string, path string) (ssh.ClientConfig, error) {
	privateKey, err := ioutil.ReadFile(path)

	if err != nil {
		return ssh.ClientConfig{}, err
	}

	signer, err := ssh.ParsePrivateKey(privateKey)

	if err != nil {
		return ssh.ClientConfig{}, err
	}

	return ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout: time.Minute,
	}, nil
}

func PrivateKeyWithPassphrase(username string, passpharase []byte, path string) (ssh.ClientConfig, error) {
	privateKey, err := ioutil.ReadFile(path)

	if err != nil {
		return ssh.ClientConfig{}, err
	}
	signer, err := ssh.ParsePrivateKeyWithPassphrase(privateKey, passpharase)

	if err != nil {
		return ssh.ClientConfig{}, err
	}

	return ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout: time.Minute,
	}, nil
}

func PasswordKey(username string, password string) (ssh.ClientConfig, error) {
	//var hostKey ssh.PublicKey
	return ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		//HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		HostKeyCallback: ssh.HostKeyCallback(func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil }),
		//HostKeyCallback: ssh.FixedHostKey(hostKey),
		Timeout: time.Minute,
	}, nil
}