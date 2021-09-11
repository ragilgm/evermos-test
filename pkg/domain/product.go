package domain

type Product struct {
	ID             uint `gorm:"primary_key"`
	ProductName    string
	Description    string
	Price          int32
	AvailableStock int32
	StockOnHold    int32
	StockSoldOut   int32
}
