package models

import (
	"gorm.io/gorm"
)

type Stuff struct {
	gorm.Model

	Name     string `gorm:"column:name"`
	LoginID  string `gorm:"column:login_id"`
	Password string `gorm:"column:password"`
}

func (Stuff) TableName() string {
	return "users"
}
