package payment

import (
	_ "gopkg.in/validator.v2"
)

type PaymentRequest struct {
	OrderId uint    `json:"order_id" validate:"required""`
	Amount  float32 `json:"amount" validate:"required"`
}
