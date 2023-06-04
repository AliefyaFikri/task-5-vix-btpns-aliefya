package database

import (
	"fmt"
	"log"
	"btpn-finpro/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectDB initializes the database connection
func ConnectDB() (*gorm.DB, error) {

	dsn := "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	return db

}

// MigrateDB performs database migrations
func MigrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
		&models.Photo{},
	)
	if err != nil {
		return fmt.Errorf("failed to perform database migrations: %w", err)
	}

	return nil
}
