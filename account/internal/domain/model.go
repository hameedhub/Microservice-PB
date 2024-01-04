package domain

import (
	"time"

	"gorm.io/gorm"
)

var (
	Active   = "Active"
	InActive = "InActive"
)

// Account model
type Account struct {
	gorm.Model
	Name           string    `json:"name"`
	Account        int64     `json:"account"`
	ClearedBalance float64   `json:"cleared_balance"`
	PendingBalance float64   `json:"pending_balance"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
