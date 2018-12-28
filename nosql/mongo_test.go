package nosql

import (
	"testing"
)

func TestServerStatus(t *testing.T) {
	status, err := ServerStatus("10.103.239.42:27017")
	if err != nil {
		t.Errorf("failed: %s", err)
	} else {
		t.Logf("%v", status.Host)
	}
}
