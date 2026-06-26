package events

type PaymentCompletedEvent struct {
	OrderID uint    `json:"order_id"`
	Amount  float64 `json:"amount"`
	Status  string  `json:"status"`
}
