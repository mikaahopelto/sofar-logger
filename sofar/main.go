package main

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	for {
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
		connStr := "amqp://guest:guest@rabbitmq:5672/"
		conn, err := amqp.Dial(connStr)
		if err != nil {
			fmt.Printf("Failed to connect to %s: %s", connStr, err.Error())
			continue
		}
		defer conn.Close()
		ch, err := conn.Channel()
		if err != nil {
			fmt.Printf("Failed to open channel. %s", err)
		}
		defer ch.Close()
	}
}
