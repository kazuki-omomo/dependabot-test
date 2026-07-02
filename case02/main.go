package main

import (
	"golang.org/x/crypto/ssh"
)

func main() {
	_, _ = ssh.Dial("tcp", "localhost:22", &ssh.ClientConfig{})
}
