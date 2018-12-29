package nosql

import (
	"testing"
)

func TestInfo(t *testing.T) {
	info, err := Info("127.0.0.1:6379", "", 0)
	if err != nil {
		t.Error("failed")
	} else {
		t.Log(info)
	}
}
