package database

import (
	"piteroni/online-store-api/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	password, err := bcrypt.GenerateFromPassword([]byte("01234567"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.Stuff{
		LoginID:  "100000",
		Password: string(password),
	}

	err = db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}
