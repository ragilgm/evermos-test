package product

import _ "gopkg.in/validator.v2"

type ProductRequest struct {
	ProductName    string `json:"product_name" validate:"required"`
	Description    string `json:"description" validate:"required"`
	Price          int32  `json:"price" validate:"required" `
	AvailableStock int32  `json:"available_stock" `
	StockOnHold    int32  `json:"stock_on_hold" `
	StockSoldOut   int32  `json:"stock_sold_out"`
}
