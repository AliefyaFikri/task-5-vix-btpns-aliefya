package database

import (
	"btpn-finpro/models"
	"gorm.io/gorm"
)

// RunMigrations runs the database migrations
func RunMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
		&models.Photo{},
	)
	if err != nil {
		return err
	}

	return nil
}
