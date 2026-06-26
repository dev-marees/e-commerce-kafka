package kafka

import (
	"context"

	kafkaGo "github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafkaGo.Writer
}

func NewProducer() *Producer {
	return &Producer{
		writer: &kafkaGo.Writer{
			Addr:  kafkaGo.TCP("localhost:9092"),
			Topic: "order-created",
		},
	}
}

func (p *Producer) Publish(msg []byte) error {
	return p.writer.WriteMessages(
		context.Background(),
		kafkaGo.Message{
			Value: msg,
		},
	)
}
