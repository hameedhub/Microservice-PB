package domain

import (
	"gorm.io/gorm"
)

type accountRepository struct {
	db *gorm.DB
}

// Account repo
type AccountRepository interface {
	Create(Account Account)
	Update(Account int64, Amount float64, Type string)
	Get(Account int64, Acc *Account)
}

func NewRepo(db *gorm.DB) *accountRepository {
	return &accountRepository{db: db}
}

// Create Transfer implements AccountRepository.
func (repo *accountRepository) Create(Acc Account) {
	repo.db.Model(&Account{}).Create(&Acc)
}

// Update implements AccountRepository.
func (repo *accountRepository) Update(Acc int64, Amount float64, Type string) {
	if Type == "credit" {
		repo.db.Model(&Account{}).Where("account = ?", Acc).Update("cleared_balance", gorm.Expr("cleared_balance + ?", Amount))
	}
	if Type == "debit" {
		repo.db.Model(&Account{}).Where("account = ?", Acc).Update("cleared_balance", gorm.Expr("cleared_balance - ?", Amount))
	}
}

// Get account information
func (repo *accountRepository) Get(Account int64, Acc *Account) {
	repo.db.Where("account =? ", Account).Find(&Acc)
}
