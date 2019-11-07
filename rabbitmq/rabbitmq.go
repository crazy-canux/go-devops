package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
	"time"
)

func Ping(url string, timeout string) bool {
	to, _ := time.ParseDuration(timeout)
	after := time.Now().Add(to)
	for {
		conn, err := amqp.Dial(url)
		if err == nil {
			defer conn.Close()
			log.Println("Rabbitmq is available.")
			return true
		}
		time.Sleep(1 * time.Second)
		if time.Now().After(after) {
			log.Println("Connection Timeout.")
			break
		}
	}
	return false
}

