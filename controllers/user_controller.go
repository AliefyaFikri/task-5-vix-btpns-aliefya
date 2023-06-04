package controllers

import (
	"btpn-finpro/app"
	"btpn-finpro/helpers"
	"btpn-finpro/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterUser handles the user registration
func RegisterUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"message": "Invalid request"})
			return
		}

		// Validate the user data
		if err := helpers.ValidateStruct(user); err != nil {
			c.JSON(400, gin.H{"message": err.Error()})
			return
		}

		// Create the user
		if err := app.CreateUser(db, &user); err != nil {
			c.JSON(500, gin.H{"message": "Failed to create user"})
			return
		}

		c.JSON(200, gin.H{"message": "User registered successfully"})
	}
}

// LoginUser handles user login
func LoginUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"message": "Invalid request"})
			return
		}

		// Retrieve the user by email
		existingUser, err := app.GetUserByEmail(db, user.Email)
		if err != nil {
			c.JSON(404, gin.H{"message": "User not found"})
			return
		}

		// Compare the provided password with the hashed password
		if !helpers.CheckPasswordHash(user.Password, existingUser.Password) {
			c.JSON(401, gin.H{"message": "Invalid credentials"})
			return
		}

		// Generate JWT token
		token, err := helpers.GenerateToken(existingUser.ID)
		if err != nil {
			c.JSON(500, gin.H{"message": "Failed to generate token"})
			return
		}

		c.JSON(200, gin.H{"token": token})
	}
}

// UpdateUser handles user update
func UpdateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("userId")

		// Get the authenticated user ID from the context
		authUserID, _ := c.Get("userID")

		// Check if the authenticated user ID matches the requested user ID
		if userID != authUserID {
			c.JSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"message": "Invalid request"})
			return
		}

		// Validate the user data
		if err := helpers.ValidateStruct(user); err != nil {
			c.JSON(400, gin.H{"message": err.Error()})
			return
		}

		// Retrieve the user from the database
		existingUser, err := app.GetUserByID(db, userID)
		if err != nil {
			c.JSON(404, gin.H{"message": "User not found"})
			return
		}

		// Update the user
		existingUser.Username = user.Username
		existingUser.Email = user.Email

		if err := app.UpdateUser(db, existingUser); err != nil {
			c.JSON(500, gin.H{"message": "Failed to update user"})
			return
		}

		c.JSON(200, gin.H{"message": "User updated successfully"})
	}
}

// DeleteUser handles user deletion
func DeleteUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("userId")

		// Get the authenticated user ID from the context
		authUserID, _ := c.Get("userID")

		// Check if the authenticated user ID matches the requested user ID
		if userID != authUserID {
			c.JSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		// Retrieve the user from the database
		existingUser, err := app.GetUserByID(db, userID)
		if err != nil {
			c.JSON(404, gin.H{"message": "User not found"})
			return
		}

		// Delete the user
		if err := app.DeleteUser(db, existingUser); err != nil {
			c.JSON(500, gin.H{"message": "Failed to delete user"})
			return
		}

		c.JSON(200, gin.H{"message": "User deleted successfully"})
	}
}
