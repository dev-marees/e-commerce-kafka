package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dev-marees/e-commerce-kafka/order-service/model"
	"github.com/dev-marees/e-commerce-kafka/order-service/service"
)

type Handler struct {
	Service *service.Service
}

func (h *Handler) Create(c *gin.Context) {

	var order model.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := h.Service.Create(&order); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, order)
}
