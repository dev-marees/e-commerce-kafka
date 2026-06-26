package producer

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer() *Producer {
	return &Producer{
		writer: &kafka.Writer{
			Addr:  kafka.TCP("kafka:9092"),
			Topic: "payment-completed",
		},
	}
}

func (p *Producer) Publish(data []byte) error {
	return p.writer.WriteMessages(
		context.Background(),
		kafka.Message{
			Value: data,
		},
	)
}
