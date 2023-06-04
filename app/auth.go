package app

import (
	"btpn-finpro/helpers"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a middleware to check if the request is authorized
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		// Validate and parse the token
		claims, err := helpers.ParseToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		// Set the user ID from the token claims to the context
		c.Set("userID", claims.UserID)

		c.Next()
	}
}
