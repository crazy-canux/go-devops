package database

import (
	"testing"
)

func TestMssqlVersion(t *testing.T) {
	version, err := MssqlVersion("127.0.0.1", 1433, "db", "user", "******")
	if err != nil {
		t.Error("failed")
	} else {
		t.Log(version)
	}
}
