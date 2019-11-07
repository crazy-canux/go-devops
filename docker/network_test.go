package docker

import (
	"testing"
)

func TestNetwork(t *testing.T) {
	id, err := CreateNetwork("lan0", "bridge", "92.168.1.0/24", "192.168.1.1")
	if err != nil {
		t.Errorf("create network error: %v", err)
	} else {
		t.Log(id)
	}
	result := RemoveNetwork("lan0")
	if result {
		t.Log("remove network pass")
	} else {
		t.Error("remove network failed.")
	}
}
