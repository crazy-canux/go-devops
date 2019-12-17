package http

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestGet(t *testing.T) {
	body, err := Get("https://10.103.238.16/capture/templates/README", "sandbox", "S0nicwall")
	if err != nil {
		t.Error("get failed.")
	}
	defer body.Close()
	content, err := ioutil.ReadAll(body)
	log.Print(string(content))
}
