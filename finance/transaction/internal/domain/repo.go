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
	UpdateTransaction(Ref int64, Status string)
	GetTransaction(Account int64, Transactions *[]Transaction)
}

func NewRepo(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db: db}
}

func (repo *transactionRepository) CreateTransaction(Trans Transaction) {
	repo.db.Model(&Transaction{}).Create(&Trans)
}

// UpdateTransaction implements TransactionRepository.
func (repo *transactionRepository) UpdateTransaction(Ref int64, Status string) {
	repo.db.Model(&Transaction{}).Where("ref = ?", Ref).Update("status", Status)
}

// GetTransaction implements TransactionRepository.
func (repo *transactionRepository) GetTransaction(Account int64, Trans *[]Transaction) {
	repo.db.Where("account = ?", Account).Find(Trans).Order("id DESC")
}
