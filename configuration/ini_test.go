package configuration

import (
	"testing"
)

var filename = "/home/canux/Src/go/src/github.com/crazy-canux/go-devops/data/configuration/grafana.ini"

func TestGetKey(t *testing.T) {
	cfg, err := LoadIni(filename)
	if err != nil {
		t.Errorf("load file error: %v", err)
	}
	appMode, err := GetKey(cfg, "", "app_mode")
	if err != nil {
		t.Errorf("get key error: %v", err)
	}
	if appMode.String() != "production" {
		t.Log(appMode)
	}
}

func TestSetKey(t *testing.T) {
	cfg, err := LoadIni(filename)
	if err != nil {
		t.Errorf("load file error: %v", err)
	}
	err = SetKey(cfg, "auth.anonymous", "enabled", "true", true)
	if err != nil {
		t.Errorf("set key error: %v", err)
	}
	if err := cfg.SaveTo(filename); err != nil {
		t.Errorf("save file error: %v", err)
	}
}
