package main

import (
	"github.com/gin-gonic/gin"

	"github.com/dev-marees/e-commerce-kafka/order-service/config"
	"github.com/dev-marees/e-commerce-kafka/order-service/handler"
	"github.com/dev-marees/e-commerce-kafka/order-service/repository"
	"github.com/dev-marees/e-commerce-kafka/order-service/routes"
	"github.com/dev-marees/e-commerce-kafka/order-service/service"

	"github.com/dev-marees/e-commerce-kafka/shared/kafka"
)

func main() {

	db := config.ConnectDB()

	repo := &repository.Repository{
		DB: db,
	}

	producer := kafka.NewProducer()

	svc := &service.Service{
		Repo:     repo,
		Producer: producer,
	}

	h := &handler.Handler{
		Service: svc,
	}

	router := gin.Default()

	routes.Register(router, h)

	router.Run(":8080")
}
