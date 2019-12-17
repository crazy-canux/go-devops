package http

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestGet(t *testing.T) {
	body, err := Get("https://127.0.0.1/path/to/file", "sandbox", "******")
	if err != nil {
		t.Error("get failed.")
	}
	defer body.Close()
	content, err := ioutil.ReadAll(body)
	log.Print(string(content))
}
