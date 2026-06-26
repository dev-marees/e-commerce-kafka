package repository

import (
	"gorm.io/gorm"

	"github.com/dev-marees/e-commerce-kafka/order-service/model"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) Create(order *model.Order) error {
	return r.DB.Create(order).Error
}
