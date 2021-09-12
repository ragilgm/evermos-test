package order

// OrderResponse line item
type OrderResponse struct {
	ID          uint   `json:"id"`
	Status      string `json:"status"`
	TotalAmount float32
	OrderItems  []OrderItemResponse `json:"order_items"`
}

// OrderItemResponse line item
type OrderItemResponse struct {
	ID       uint         `json:"id"`
	OrderID  uint         `json:"order_id"`
	ItemID   uint         `json:"item_id"`
	Item     ItemResponse `json:"item"`
	Quantity int          `json:"quantity"`
	Total    float32      `json:"total"`
}

// ItemResponse line item
type ItemResponse struct {
	ID       uint    `json:"id"`
	ItemName string  `json:"item_name"`
	Amount   float32 `json:"amount"`
}
