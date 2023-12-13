package main

import (
	"broker/client"
	"broker/event"
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const webPort = "80"

type Config struct {
	AuthClient   client.AuthClient
	LoggerClient client.LoggerClient
	MailerClient client.MailerClient
	LogRPCClient client.LogRPCClient
	LogProducer  event.Producer
}

func main() {
	conn, err := connect()
	if err != nil {
		log.Panic("Failed to connect to rabbit mq", err)
	}
	defer conn.Close()
	producer, err := event.NewProducer(conn)
	if err != nil {
		log.Panic("Failed to connect to rabbit mq", err)
	}
	app := Config{
		AuthClient:   client.NewAuthClient(),
		LoggerClient: client.NewLoggerClient(),
		MailerClient: client.NewMailerClient(),
		LogRPCClient: client.NewLogRPCClient(),
		LogProducer:  producer,
	}
	log.Printf("Starting broker service on port %s \n", webPort)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
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
