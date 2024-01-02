package domain

import (
	"gorm.io/gorm"
)

type transferRepository struct {
	db *gorm.DB
}

// Transfer repo
type TransferRepository interface {
	Create(Amount int64)
	Update(ID int64, Status string)
}

func NewRepo(db *gorm.DB) *transferRepository {
	return &transferRepository{db: db}
}

// Create Transfer implements TransferRepository.
func (repo *transferRepository) Create(Amount int64) {
	// repo.db.Model(&Transfer{}).Create(&Trans)
}

// Update implements TransferRepository.
func (repo *transferRepository) Update(ID int64, Status string) {
	repo.db.Where("id = ?", ID).Update("Status", Status)
}
