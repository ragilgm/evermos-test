package checkout

import (
	_ "gopkg.in/validator.v2"
)

type CheckOutRequest struct {
	OrderId uint  `json:"order_id" validate:"required"`
	Amount  int32 `json:"amount" validate:"required"`
}
