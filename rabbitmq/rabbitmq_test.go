package rabbitmq

import (
	"testing"
)

func TestPing(t *testing.T) {
	result := Ping("amqp://sandbox:password@10.103.239.61:5672/", "10s")
	if result {
		t.Log("pass.")
	} else {
		t.Error("failed.")
	}
}