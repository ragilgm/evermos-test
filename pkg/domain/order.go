package domain

// Order
type Order struct {
	ID          uint `gorm:"primary_key"`
	Status      string
	TotalAmount int32
	OrderItems  []OrderItem
}

// Order line item
type OrderItem struct {
	ID       uint `gorm:"primary_key"`
	OrderID  uint
	ItemID   uint
	Item     Item
	Quantity int32
	Total    int32
}

// Product
type Item struct {
	ID       uint `gorm:"primary_key"`
	ItemName string
	Amount   int32
}
