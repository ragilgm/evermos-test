package checkout

import "time"

type CheckoutResponse struct {
	ID                uint      `json:"id"`
	TransactionNumber string    `json:"transaction_number"`
	OrderId           uint      `json:"order_id"`
	TransactionDate   time.Time `json:"transaction_date"`
	ExpiredAt         time.Time `json:"expired_at"`
	TransactionToken  string    `json:"transaction_token"`
	Status            string    `json:"status"`
	Amount            float32   `json:"amount"`
}
