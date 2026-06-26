package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/dev-marees/e-commerce-kafka/order-service/model"
)

func ConnectDB() *gorm.DB {

	dsn := "host=postgres user=postgres password=postgres dbname=ecommerce port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Order{})

	return db
}
