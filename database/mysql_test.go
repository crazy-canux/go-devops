package database

import (
	"testing"
)

func TestMysqlVersion(t *testing.T) {
	version, err := MysqlVersion("127.0.0.1", 3306, "sandbox", "sandbox", "password")
	if err != nil {
		t.Errorf("failed: %s", err)
	} else {
		t.Log(version)
	}
}
