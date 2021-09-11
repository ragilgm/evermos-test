package payment

import "time"

type PaymentResponse struct {
	ID                uint      `json:"id"`
	TransactionNumber string    `json:"transaction_number"`
	OrderId           uint      `json:"order_id"`
	TransactionDate   time.Time `json:"transaction_date"`
	Status            string    `json:"status"`
	Amount            float32   `json:"amount"`
}
