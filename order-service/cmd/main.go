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

	cfg := config.LoadConfig()

	db := config.ConnectDB(cfg)

	producer := kafka.NewProducer(cfg)

	repo := &repository.Repository{
		DB: db,
	}

	svc := &service.Service{
		Repo:     repo,
		Producer: producer,
	}

	h := &handler.Handler{
		Service: svc,
	}

	router := gin.Default()

	routes.Register(router, h)

	router.Run(":" + cfg.AppPort)
}
