package order

type OrderRequest struct {
	OrderItems []OrderItemRequest `json:"order_items"`
}

// Order line item
type OrderItemRequest struct {
	ItemID   uint  `json:"item_id"`
	Quantity int32 `json:"quantity"`
}
