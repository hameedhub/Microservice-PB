package domain

import (
	"gorm.io/gorm"
)

var (
	Credit = "credit"
	Debit  = "debit"
)

var (
	Success = "Successful"
	Failed  = "Failed"
	Pending = "Pending"
)

// Transaction model
type Transaction struct {
	gorm.Model
	Amount  float64 `json:"amount"`
	Account int64   `json:"account"`
	Type    string  `json:"type"`
	Status  string  `json:"status"`
}
