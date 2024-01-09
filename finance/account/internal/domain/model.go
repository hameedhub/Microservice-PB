package domain

import (
	"time"
)

var (
	Success = "Successful"
	Failed  = "Failed"
	Pending = "Pending"
)

// status
var (
	Active   = "Active"
	InActive = "InActive"
)

// Account model
type Account struct {
	ID             int64     `json:"-"`
	Name           string    `json:"name"`
	Account        int64     `json:"account"`
	ClearedBalance float64   `json:"cleared_balance"`
	PendingBalance float64   `json:"pending_balance"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Deposit struct {
	Account int64   `json:"account"`
	Amount  float64 `json:"amount"`
	Type    string  `json:"type"`
}

type Transfer struct {
	CreditAccount int64   `json:"credit_account"`
	DebitAccount  int64   `json:"debit_account"`
	Amount        float64 `json:"amount"`
	Ref           int64   `json:"ref"`
	Status        string  `json:"status"`
}
