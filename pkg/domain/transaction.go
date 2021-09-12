package domain

import "time"

const (
	Pending = "PENDING"
	Success = "SUCCESS"
	Failed  = "FAILED"
)

type Transaction struct {
	ID                uint `gorm:"primary_key"`
	TransactionNumber string
	OrderId           uint
	TransactionDate   time.Time `gorm:"autoCreateTime"`
	Status            string
	ExpiredAt         time.Time
	TransactionToken  string
	Amount            int32
}
