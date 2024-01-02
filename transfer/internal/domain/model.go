package domain

import (
	"time"

	"gorm.io/gorm"
)

var (
	Success = "Sent"
	Failed  = "Failed"
	Pending = "Pending"
)

// Transfer model
type Transfer struct {
	gorm.Model
	CreditAccount int64     `json:"credit_account"`
	DebitAccount  int64     `json:"debit_account"`
	Amount        float64   `json:"amount"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
