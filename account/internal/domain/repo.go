package domain

import (
	"gorm.io/gorm"
)

type accountRepository struct {
	db *gorm.DB
}

// Account repo
type AccountRepository interface {
	Create(Amount int64)
	Update(ID int64, Status string)
}

func NewRepo(db *gorm.DB) *accountRepository {
	return &accountRepository{db: db}
}

// Create Transfer implements AccountRepository.
func (repo *accountRepository) Create(Amount int64) {
	// repo.db.Model(&Transfer{}).Create(&Trans)
}

// Update implements AccountRepository.
func (repo *accountRepository) Update(ID int64, Status string) {
	repo.db.Where("id = ?", ID).Update("Status", Status)
}
