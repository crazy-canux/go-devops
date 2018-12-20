package tickstack

import (
	"testing"
)

func TestGetSmtp(t *testing.T) {
	resp, err := GetSmtp("10.103.239.60", "9092")
	if err != nil {
		t.Error("test getsmtp failed.")
	} else {
		t.Log(resp)
	}
}