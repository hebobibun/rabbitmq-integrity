package main

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	amqpUri := "amqp://username:password@localhost:5672/"
	// Establish connection to RabbitMQ server
	conn, err := amqp.Dial(amqpUri)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	log.Println("rabbitMQ publisher is connected")

	// Create a channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declare a queue
	q, err := ch.QueueDeclare(
		"hello", // queue name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	counter := 1

	// Publish messages periodically
	for {
		message := "Hello, RabbitMQ!"
		body := fmt.Sprintf(" [%v] Sent %s", counter, message)
		err := ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		failOnError(err, "Failed to publish a message")

		log.Println(body)
		counter++

		time.Sleep(5 * time.Second) // adjust sleep interval as needed
	}
}
