package kafka

import (
	"context"

	"github.com/dev-marees/e-commerce-kafka/order-service/config"
	"github.com/segmentio/kafka-go"
	kafkaGo "github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafkaGo.Writer
}

func NewProducer(cfg *config.Config) *Producer {
	return &Producer{
		writer: &kafka.Writer{
			Addr:  kafka.TCP(cfg.KafkaBroker),
			Topic: cfg.KafkaTopic,
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
