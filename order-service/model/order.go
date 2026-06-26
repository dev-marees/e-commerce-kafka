package model

type Order struct {
	ID        uint `gorm:"primaryKey"`
	ProductID uint
	Quantity  int
	Amount    float64
}
