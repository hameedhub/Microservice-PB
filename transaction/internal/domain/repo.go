package domain

import "context"

// Transaction repo
type TransactionRepository interface {
	CreateTransaction(ctx context.Context, Amount float64, Account, Ref int64, Status string)
	GetTransaction(ctx context.Context, Account int64)
}
