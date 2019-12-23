package ssh


import (
	"fmt"
	"os"
	"testing"
	"time"
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

	o, e, err := client.Run("pwd" )
	if err != nil {
		t.Error("run failed")
	}
	fmt.Print(o, e)
}

func TestNewClientWithTimeout(t *testing.T) {
	clientConfig, _ := PasswordKey("canux", "******")
	client := NewClientWithTimeout("127.0.0.1:22", &clientConfig, time.Second * 3)
	defer client.Close()

	err := client.Connect()
	if err != nil {
		t.Error("connection failed")
	}

	o, e, err := client.Run("sleep 2" )
	if err != nil {
		t.Errorf("run failed: %s", err.Error())
	}
	fmt.Print(o, e)
}

func TestNewClientWithBasicAuth(t *testing.T) {
	client := NewClientWithBasicAuth("127.0.0.1:22", "canux", "******")
	defer client.Close()

	err := client.Connect()
	if err != nil {
		t.Errorf("connection failed: %s", err.Error())
	}

	o, e, err := client.Run("pwd" )
	if err != nil {
		t.Errorf("run failed: %s", err.Error())
	}
	fmt.Print(o, e)
}

func TestNewClientWithBasicAuthAndTimeout(t *testing.T) {
	client := NewClientWithBasicAuthAndTimeout("127.0.0.1:22", "canux", "******", time.Second * 3)
	defer client.Close()

	err := client.Connect()
	if err != nil {
		t.Errorf("connection failed: %s", err.Error())
	}

	o, e, err := client.Run("pwd")
	if err != nil {
		t.Errorf("run failed: %s", err.Error())
	}
	fmt.Print(o, e)
}

func TestRunWithSudo(t *testing.T) {
	client := NewClientWithBasicAuthAndTimeout("127.0.0.1:22", "canux", "******", time.Second * 3)
	defer client.Close()

	err := client.Connect()
	if err != nil {
		t.Errorf("connection failed: %s", err.Error())
	}

	o, e, err := client.RunWithSudo("sudo echo 1")
	if err != nil {
		t.Errorf("run with sudo failed: %s", err.Error())
	}
	fmt.Print(o, e)
}
