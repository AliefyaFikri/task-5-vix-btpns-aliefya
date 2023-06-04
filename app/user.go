package app

import (
	"btpn-finpro/models"
	"gorm.io/gorm"
)

// CreateUser creates a new user
func CreateUser(db *gorm.DB, user *models.User) error {
	return db.Create(user).Error
}

// GetUserByEmail retrieves a user by email
func GetUserByEmail(db *gorm.DB, email string) (*models.User, error) {
	var user models.User
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates a user
func UpdateUser(db *gorm.DB, user *models.User) error {
	return db.Save(user).Error
}

// DeleteUser deletes a user
func DeleteUser(db *gorm.DB, user *models.User) error {
	return db.Delete(user).Error
}
