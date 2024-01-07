package domain

import (
	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

// Transaction repo
type TransactionRepository interface {
	CreateTransaction(Trans Transaction)
	UpdateTransaction(Account int64, Status string, Transaction *Transaction)
	GetTransaction(Account int64, Transactions *[]Transaction)
}

func NewRepo(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db: db}
}

func (repo *transactionRepository) CreateTransaction(Trans Transaction) {
	repo.db.Model(&Transaction{}).Create(&Trans)
}

// UpdateTransaction implements TransactionRepository.
func (repo *transactionRepository) UpdateTransaction(Account int64, Status string, Tran *Transaction) {
	repo.db.Model(&Transaction{}).Where("account = ?", Account).Update("status", Status).Find(Tran)
}

// GetTransaction implements TransactionRepository.
func (repo *transactionRepository) GetTransaction(Account int64, Trans *[]Transaction) {
	repo.db.Where("account = ?", Account).Limit(50).Find(Trans)
}
