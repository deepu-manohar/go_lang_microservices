package event

import (
	"encoding/json"
	"fmt"
	"listener-service/client"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn      *amqp.Connection
	queueName string
}

type Payload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func NewConsumer(conn *amqp.Connection) (Consumer, error) {
	consumer := Consumer{
		conn: conn,
	}
	err := consumer.setup()
	if err != nil {
		log.Println("Failed to create channel")
		log.Panic(err)
		return Consumer{}, err
	}
	return consumer, nil
}

func (consumer *Consumer) setup() error {
	channel, err := consumer.conn.Channel()
	if err != nil {
		return err
	}
	return declareExchange(channel)
}

func (c *Consumer) Listen(topics []string) error {
	channel, err := c.conn.Channel()
	if err != nil {
		log.Println("Error getting channels for listening ", err)
		return err
	}
	defer channel.Close()
	queue, err := declareRandomQueue(channel)

	if err != nil {
		log.Println("Error getting queue ", err)
		return err
	}

	for _, topic := range topics {
		err = channel.QueueBind(
			queue.Name,
			topic,
			"logs_topic",
			false,
			nil,
		)
		if err != nil {
			log.Println("Error binding queue ", err)
			return err
		}
	}

	messages, err := channel.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Println("Error getting queue ", err)
		return err
	}

	forever := make(chan bool)

	go func() {
		for message := range messages {
			var payload Payload
			_ = json.Unmarshal(message.Body, &payload)
			log.Printf("Got message %s, %s", payload.Name, payload.Data)
			go handlePayload(payload)
		}
	}()

	fmt.Printf("Waiting for message [Exchange, Queue] [logs_topic, %s]\n", queue.Name)
	<-forever
	return nil
}

func handlePayload(payload Payload) {
	switch payload.Name {
	case "log", "event":
		err := logEvent(payload)
		if err != nil {
			log.Println("Error consuming logevent ", err)
		}
	case "auth":
	default:
		err := logEvent(payload)
		if err != nil {
			log.Println("Error consuming logevent ", err)
		}
	}
}

func logEvent(payload Payload) error {
	logClient := client.NewLoggerClient()
	var logRequest = client.LogRequest{}
	json.Unmarshal([]byte(payload.Data), &logRequest)
	logClient.Log(logRequest)
	return nil
}
