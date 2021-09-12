package order

// OrderRequest line item
type OrderRequest struct {
	OrderItems []OrderItemRequest `json:"order_items"`
}

// OrderItemRequest line item
type OrderItemRequest struct {
	ItemID   uint  `json:"item_id"`
	Quantity int32 `json:"quantity"`
}
