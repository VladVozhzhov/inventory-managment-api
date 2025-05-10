package models

import (
	"github.com/VladVozhzhov/inventory-managment-api/utils"
	"gorm.io/gorm"
)

type Stock struct {
	ID        string `gorm:"primaryKey;size:20"`
	ProductID string
	Change    int
	Reason    string
	UserID    string
}

func (s *Stock) BeforeCreate(tx *gorm.DB) (err error) {
	if s.ID == "" {
		s.ID = utils.GenerateRandomID()
	}
	return
}
