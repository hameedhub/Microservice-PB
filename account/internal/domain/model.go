package domain

import (
	"time"
)

// Domain producer
var (
	CreateAccount  = "create_account"
	AccountDeposit = "account_deposit"
)

// Domain consumer

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
