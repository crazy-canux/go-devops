package nosql

import (
	"testing"
)

func TestInfo(t *testing.T) {
	info, err := Info("10.103.239.75:6379", "", 0)
	if err != nil {
		t.Error("failed")
	} else {
		t.Log(info)
	}
}
