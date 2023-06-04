package middlewares

import (
	"btpn-finpro/helpers"
	"github.com/gin-gonic/gin"
)

// AuthenticateUser is a middleware to authenticate the user
func AuthenticateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := helpers.GetBearerToken(c.Request)

		if tokenString == "" {
			c.JSON(401, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		claims, err := helpers.ParseToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}
