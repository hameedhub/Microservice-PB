package domain

import (
	"fmt"

	"gorm.io/gorm"
)

type transferRepository struct {
	db *gorm.DB
}

// Transfer repo
type TransferRepository interface {
	Create(Trans *Transfer)
	Update(ID int64, Status string)
}

func NewRepo(db *gorm.DB) *transferRepository {
	return &transferRepository{db: db}
}

// Create Transfer implements TransferRepository.
func (repo *transferRepository) Create(Trans *Transfer) {
	repo.db.Model(&Transfer{}).Create(&Trans)
}

// Update implements TransferRepository.
func (repo *transferRepository) Update(Ref int64, Status string) {
	err := repo.db.Where("ref = ?", Ref).Update("status", Status).Error
	if err != nil {
		fmt.Println(err)
	}
}
