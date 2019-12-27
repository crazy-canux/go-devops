package database

import (
	"testing"
)

func TestMssqlVersion(t *testing.T) {
	version, err := MssqlVersion("10.103.239.70", 1433, "sandbox", "sandbox", "P@ssword")
	if err != nil {
		t.Error("failed")
	} else {
		t.Log(version)
	}
}
