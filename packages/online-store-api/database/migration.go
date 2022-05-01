package database

import (
	"piteroni/online-store-api/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	m := []interface{}{
		&models.Stuff{},
	}

	for _, model := range m {
		err := db.AutoMigrate(model)
		if err != nil {
			return err
		}
	}

	return nil
}

func DropTables(db *gorm.DB) error {
	return db.Migrator().DropTable(
		&models.Stuff{},
	)
}
