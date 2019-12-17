package http

import (
	"testing"
)

func TestDownload(t *testing.T) {
	result := Download("https://10.103.238.16/capture/templates/templates_4/Sandbox-WIN7-AMD64/base.img",
		"/home/canux/Src/go/src/github.com/crazy-canux/go-devops/base.img",
		"sandbox", "S0nicwall")
	if result {
		t.Log("yes")
	} else {
		t.Error("no")
	}
}
