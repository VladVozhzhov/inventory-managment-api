package models

import (
	"github.com/VladVozhzhov/inventory-managment-api/utils"
	"gorm.io/gorm"
)

type User struct {
	ID       string `gorm:"primaryKey;size:20"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"default:'staff';not null"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = utils.GenerateRandomID()
	}
	return
}
