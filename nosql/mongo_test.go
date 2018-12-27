package nosql

import (
	"testing"
)

func TestServerStatus(t *testing.T) {
	ServerStatus("10.103.239.43:27017")
}
