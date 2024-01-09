package domain

import (
	"gorm.io/gorm"
)

var (
	Success = "Successful"
	Failed  = "Failed"
	Pending = "Pending"
)

// Transfer model
type Transfer struct {
	gorm.Model
	CreditAccount int64   `json:"credit_account"`
	DebitAccount  int64   `json:"debit_account"`
	Amount        float64 `json:"amount"`
	Ref           int64   `json:"ref"`
	Status        string  `json:"status"`
}
