package database

import (
	"testing"
)

func TestMysqlVersion(t *testing.T) {
	version, err := MysqlVersion("10.103.239.75", 3306, "sandbox", "sandbox", "password")
	if err != nil {
		t.Errorf("failed: %s", err)
	} else {
		t.Log(version)
	}
}
