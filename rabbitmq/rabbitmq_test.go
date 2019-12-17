package rabbitmq

import (
	"testing"
)

func TestPing(t *testing.T) {
	result := Ping("amqp://user:******@127.0.0.1:5672/", "10s")
	if result {
		t.Log("pass.")
	} else {
		t.Error("failed.")
	}
}
