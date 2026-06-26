package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/dev-marees/e-commerce-kafka/order-service/handler"
)

func Register(router *gin.Engine, h *handler.Handler) {

	router.POST("/orders", h.Create)
}
