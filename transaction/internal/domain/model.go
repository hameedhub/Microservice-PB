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
	ID        int64
	Amount    float64
	Account   int64
	Ref       int64
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
