package configuration

import (
	"fmt"
	"testing"
)

var fileName = "/home/canux/Src/go/src/github.com/crazy-canux/go-devops/data/configuration/test.toml"

type Constraint struct {
	name string
	version string
}

func TestParseFromFile(t *testing.T) {
    var config Constraint

	err := ParseFromFile(fileName, &config)
	if err != nil {
		t.Errorf("parse from file failed: %v", err)
	}
	t.Log(config)
	fmt.Print(config)
}
