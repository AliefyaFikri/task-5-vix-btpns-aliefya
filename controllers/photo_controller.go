package controllers

import (
	"btpn-finpro/app"
	"btpn-finpro/helpers"
	"btpn-finpro/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreatePhoto handles photo creation
func CreatePhoto(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var photo models.Photo
		if err := c.ShouldBindJSON(&photo); err != nil {
			c.JSON(400, gin.H{"message": "Invalid request"})
			return
		}

		// Get the authenticated user ID from the context
		userID, _ := c.Get("userID")

		// Assign the user ID to the photo
		photo.UserID = userID.(string)

		// Create the photo
		if err := app.CreatePhoto(db, &photo); err != nil {
			c.JSON(500, gin.H{"message": "Failed to create photo"})
			return
		}

		c.JSON(200, gin.H{"message": "Photo created successfully"})
	}
}

// GetPhotos retrieves all photos
func GetPhotos(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var photos []models.Photo

		// Retrieve all photos from the database
		if err := db.Find(&photos).Error; err != nil {
			c.JSON(500, gin.H{"message": "Failed to retrieve photos"})
			return
		}

		c.JSON(200, photos)
	}
}

// UpdatePhoto handles photo update
func UpdatePhoto(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		photoID := c.Param("photoId")

		// Get the authenticated user ID from the context
		authUserID, _ := c.Get("userID")

		// Check if the authenticated user owns the photo
		if !app.IsUserPhotoOwner(db, authUserID.(string), photoID) {
			c.JSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		var photo models.Photo
		if err := c.ShouldBindJSON(&photo); err != nil {
			c.JSON(400, gin.H{"message": "Invalid request"})
			return
		}

		// Retrieve the photo from the database
		existingPhoto, err := app.GetPhotoByID(db, photoID)
		if err != nil {
			c.JSON(404, gin.H{"message": "Photo not found"})
			return
		}

		// Update the photo
		existingPhoto.Title = photo.Title
		existingPhoto.Caption = photo.Caption
		existingPhoto.PhotoURL = photo.PhotoURL

		if err := app.UpdatePhoto(db, existingPhoto); err != nil {
			c.JSON(500, gin.H{"message": "Failed to update photo"})
			return
		}

		c.JSON(200, gin.H{"message": "Photo updated successfully"})
	}
}

// DeletePhoto handles photo deletion
func DeletePhoto(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		photoID := c.Param("photoId")

		// Get the authenticated user ID from the context
		authUserID, _ := c.Get("userID")

		// Check if the authenticated user owns the photo
		if !app.IsUserPhotoOwner(db, authUserID.(string), photoID) {
			c.JSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		// Retrieve the photo from the database
		existingPhoto, err := app.GetPhotoByID(db, photoID)
		if err != nil {
			c.JSON(404, gin.H{"message": "Photo not found"})
			return
		}

		// Delete the photo
		if err := app.DeletePhoto(db, existingPhoto); err != nil {
			c.JSON(500, gin.H{"message": "Failed to delete photo"})
			return
		}

		c.JSON(200, gin.H{"message": "Photo deleted successfully"})
	}
}
