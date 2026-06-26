package main

import (
	"log"

	"github.com/dev-marees/e-commerce-kafka/payment-service/consumer"
	"github.com/dev-marees/e-commerce-kafka/payment-service/producer"
	"github.com/dev-marees/e-commerce-kafka/payment-service/service"
)

func main() {

	log.Println("Payment Service Started")

	producer := producer.NewProducer()

	svc := &service.Service{
		Producer: producer,
	}

	consumer.Start(svc)
}
