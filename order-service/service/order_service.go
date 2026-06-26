package service

import (
	"encoding/json"

	"github.com/dev-marees/e-commerce-kafka/order-service/model"
	"github.com/dev-marees/e-commerce-kafka/order-service/repository"

	"github.com/dev-marees/e-commerce-kafka/shared/events"
	"github.com/dev-marees/e-commerce-kafka/shared/kafka"
)

type Service struct {
	Repo     *repository.Repository
	Producer *kafka.Producer
}

func (s *Service) Create(order *model.Order) error {

	if err := s.Repo.Create(order); err != nil {
		return err
	}

	event := events.OrderCreatedEvent{
		OrderID:   order.ID,
		ProductID: order.ProductID,
		Quantity:  order.Quantity,
		Amount:    order.Amount,
	}

	data, _ := json.Marshal(event)

	return s.Producer.Publish(data)
}
