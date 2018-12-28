package ssh

import (
	"golang.org/x/crypto/ssh"
	"log"
	"bytes"
)

// Run command on remote server by ssh.
func Run(ip, user, pw, cmd string) (string, string, error) {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pw),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", ip, config)
	if err != nil {
		log.Println("Failed to dial.")
		return "", "", err
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Println("Failed to create session.")
		return "", "", err
	}
	defer session.Close()

	var stdOut, stdErr bytes.Buffer
	session.Stdout = &stdOut
	session.Stderr = &stdErr

	err = session.Run(cmd)
	if err != nil {
		log.Println("Run command failed.")
		return "", "", err
	}
	return stdOut.String(), stdErr.String(), nil
}
