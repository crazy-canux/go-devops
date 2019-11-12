package tickstack

import (
	"testing"
)

func TestGetSmtp(t *testing.T) {
	resp, err := GetSmtp("127.0.0.1", "9092")
	if err != nil {
		t.Error("test getsmtp failed.")
	} else {
		t.Log(resp)
	}
}
