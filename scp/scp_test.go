package scp


import (
	"os"
	"testing"
)

func TestScp(t *testing.T) {
	clientConfig, _ := PasswordKey("canux", "S0nicwall")
	client := NewClient("10.103.239.40:22", &clientConfig)
	defer client.Close()

	err := client.Connect()
	if err != nil {
		t.Error("failed")
	}

	f, _ := os.Open("/home/canux/Src/go/src/github.com/crazy-canux/go-devops/data/cryptographic/APT1.yar")
	defer f.Close()

	err = client.CopyFile(f, "/home/canux/apt.yar", "0655")
	if err != nil {
		t.Error("copyfile failed")
	}
}