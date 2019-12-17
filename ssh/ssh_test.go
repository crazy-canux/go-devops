package ssh


import (
	"fmt"
	"os"
	"testing"
)

func TestScp(t *testing.T) {
	clientConfig, _ := PasswordKey("canux", "******")
	client := NewClient("127.0.0.1:22", &clientConfig)
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

func TestNewClient(t *testing.T) {
	clientConfig, _ := PasswordKey("canux", "******")
	client := NewClient("127.0.0.1:22", &clientConfig)
	defer client.Close()

	err := client.Connect()
	if err != nil {
		t.Error("connection failed")
	}

	o, e, err := client.Run("pwd")
	if err != nil {
		t.Error("run failed")
	}
	fmt.Print(o, e)
}

func TestNewClientWithBasicAuth(t *testing.T) {
	client := NewClientWithBasicAuth("127.0.0.1:22", "canux", "******")
	defer client.Close()

	err := client.Connect()
	if err != nil {
		t.Error("connection failed")
	}

	o, e, err := client.Run("pwd")
	if err != nil {
		t.Error("run failed")
	}
	fmt.Print(o, e)
}