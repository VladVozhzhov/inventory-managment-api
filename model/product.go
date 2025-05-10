package models

import (
	"github.com/VladVozhzhov/inventory-managment-api/utils"
	"gorm.io/gorm"
)

type Product struct {
	ID          string `gorm:"primaryKey;size:20"`
	Name        string `gorm:"not null"`
	SKU         string `gorm:"uniqueIndex;not null"`
	Category    string
	Quantity    int `gorm:"not null;default:0"`
	Supplier    string
	Description string
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == "" {
		p.ID = utils.GenerateRandomID()
	}
	return
}
