package domain

import (
	"time"

	"gorm.io/gorm"
)

var (
	Success = "Successful"
	Failed  = "Failed"
	Pending = "Pending"
)

// Transaction model
type Transaction struct {
	gorm.Model
	Amount    float64   `json:"amount"`
	Account   int64     `json:"account"`
	Ref       int64     `json:"ref"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
