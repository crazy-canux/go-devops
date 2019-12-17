package http

import (
	"testing"
)

func TestDownload(t *testing.T) {
	result := Download("https://127.0.0.1/path/to/file",
		"/path/to/file",
		"sandbox", "******")
	if result {
		t.Log("yes")
	} else {
		t.Error("no")
	}
}
