package nosql

import (
	"testing"
)

func TestServerStatus(t *testing.T) {
	status, err := ServerStatus("127.0.0.1:27017")
	if err != nil {
		t.Errorf("failed: %s", err)
	} else {
		t.Logf("%v", status.Host)
	}
}
