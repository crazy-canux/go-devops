package ssh

import (
	"testing"
)

func TestRun(t *testing.T) {
	output, error, err := Run("localhost:22", "canux", "canux", "who")
	if err != nil {
		t.Error("failed")
	} else {
		t.Log(output)
		t.Log(error)
	}
}
