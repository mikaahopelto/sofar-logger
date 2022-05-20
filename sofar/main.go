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
			continue
		}
		defer ch.Close()
		// We create a Queue to send the message to.
		q, err := ch.QueueDeclare(
			"sofar-queue", // name
			false,         // durable
			false,         // delete when unused
			false,         // exclusive
			false,         // no-wait
			nil,           // arguments
		)
		if err != nil {
			fmt.Printf("Failed to declare queue. %s", err)
			continue
		}
		body := fmt.Sprintf(`{"Testing": "%s"}`, time.Now().UTC().String())
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        []byte(body),
			})
		if err != nil {
			fmt.Printf("Failed to send message. %s", err)
			continue
		}
	}
}
