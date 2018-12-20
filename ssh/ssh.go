package main

import (
	"golang.org/x/crypto/ssh"
	"log"
	"bytes"
	"fmt"
)

func main() {
	config := &ssh.ClientConfig{
		User: "canux",
		Auth: []ssh.AuthMethod{
			ssh.Password("canux"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", "127.0.0.1:22", config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	var stdOut, stdErr bytes.Buffer
	session.Stdout = &stdOut
	session.Stderr = &stdErr

	session.Run("who")

	fmt.Println(stdOut.String())
	fmt.Println(stdErr.String())
}
