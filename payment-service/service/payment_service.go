package service

import (
	"encoding/json"
	"log"
	"time"

	"github.com/dev-marees/e-commerce-kafka/payment-service/producer"
	"github.com/dev-marees/e-commerce-kafka/shared/events"
)

type Service struct {
	Producer *producer.Producer
}

func (s *Service) Process(order events.OrderCreatedEvent) error {

	log.Printf("Processing payment for order %d", order.OrderID)

	time.Sleep(2 * time.Second)

	event := events.PaymentCompletedEvent{
		OrderID: order.OrderID,
		Amount:  order.Amount,
		Status:  "SUCCESS",
	}

	data, _ := json.Marshal(event)

	return s.Producer.Publish(data)
}
