package main

import (
	"errors"
	"listener-service/event"
	"log"
	"math"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := connect()
	if err != nil {
		log.Panic("Failed to connect to rabbit mq", err)
	}
	defer conn.Close()
	consumer, err := event.NewConsumer(conn)
	if err != nil {
		panic(err)
	}
	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})
	if err != nil {
		log.Panic("Failed to liste ", err)
	}
}

func connect() (*amqp.Connection, error) {
	var retries = 15
	var backoff = 1 * time.Second

	for i := 0; i < retries; i++ {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			log.Println("rabbitmq is not ready")
		} else {
			log.Println("rabbitmq is ready")
			return c, nil
		}
		backoff = time.Duration(math.Pow(float64(i), 2)) * time.Second
		log.Println("Backing off for duration .. ", backoff)
		time.Sleep(backoff)
	}
	return nil, errors.New("couldn't connect to rabbitmq")
}
