package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	for {
		time.Sleep(5 * time.Second)
		connStr := fmt.Sprintf("amqp://%s:%s@%s/", os.Getenv("AMQP_USER"), os.Getenv("AMQP_PASSWORD"), os.Getenv("AMQP_HOST"))
		conn, err := amqp.Dial(connStr)
		if err != nil {
			log.Printf("Failed to connect to %s: %s", connStr, err.Error())
			continue
		}
		defer conn.Close()
		ch, err := conn.Channel()
		if err != nil {
			log.Printf("Failed to open channel. %s", err)
			continue
		}
		defer ch.Close()
		// We create a Queue to send the message to.
		q, err := ch.QueueDeclare(
			os.Getenv("AMQP_QUEUE"), // name
			false,                   // durable
			false,                   // delete when unused
			false,                   // exclusive
			false,                   // no-wait
			nil,                     // arguments
		)
		if err != nil {
			log.Printf("Failed to declare queue. %s", err)
			continue
		}
		body := fmt.Sprintf(`{"temperature": 3, "b": 4, "timestamp": %d}`, time.Now().UnixMilli())
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
			log.Printf("Failed to send message. %s", err)
			continue
		}
		fmt.Println(body)
	}
}
