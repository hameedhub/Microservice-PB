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
	Ref     int64   `json:"ref"`
	Status  string  `json:"status"`
}

type Transfer struct {
	CreditAccount int64   `json:"credit_account"`
	DebitAccount  int64   `json:"debit_account"`
	Amount        float64 `json:"amount"`
	Ref           int64   `json:"ref"`
	Status        string  `json:"status"`
}
