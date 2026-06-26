package consumer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"

	"github.com/dev-marees/e-commerce-kafka/payment-service/service"
	"github.com/dev-marees/e-commerce-kafka/shared/events"
)

func Start(svc *service.Service) {

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   "order-created",
		GroupID: "payment-group",
	})

	for {

		msg, err := reader.ReadMessage(context.Background())

		if err != nil {
			log.Println(err)
			continue
		}

		var event events.OrderCreatedEvent

		json.Unmarshal(msg.Value, &event)

		svc.Process(event)
	}
}
