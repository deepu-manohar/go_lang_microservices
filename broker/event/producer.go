package event

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	conn *amqp.Connection
}

func NewProducer(conn *amqp.Connection) (Producer, error) {
	producer := Producer{
		conn: conn,
	}
	err := producer.Setup()
	if err != nil {
		return Producer{}, err
	}
	return producer, nil
}

func (p *Producer) Setup() error {
	channel, err := p.conn.Channel()
	if err != nil {
		log.Panic("failed to get channel", err)
		return err
	}
	defer channel.Close()
	err = DeclareExchange(channel)
	if err != nil {
		log.Panic("failed to get exchange", err)
		return err
	}
	return nil
}

func DeclareExchange(channel *amqp.Channel) error {
	return channel.ExchangeDeclare(
		"log_topic",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
}
func (p *Producer) Produce(event string, severity string) error {
	channel, err := p.conn.Channel()
	if err != nil {
		log.Panic("failed to get channel", err)
		return err
	}
	defer channel.Close()
	log.Printf("producing to channel %s, %s\n", event, severity)
	err = channel.Publish("logs_topic", severity, false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(event),
		})
	if err != nil {
		log.Panic("failed to get channel", err)
	}
	return nil
}
