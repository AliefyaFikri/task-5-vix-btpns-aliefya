package router

import (
	"btpn-finpro/app"
	"btpn-finpro/controllers"
	"btpn-finpro/middlewares"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the Gin router
func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Initialize the app
	app := app.NewApp(db)

	// Public routes
	r.POST("/users/register", controllers.RegisterUser(app))
	r.POST("/users/login", controllers.LoginUser(app))

	// Private routes
	private := r.Group("/api")
	private.Use(middlewares.AuthenticateUser())

	private.PUT("/users/:userID", controllers.UpdateUser(app))
	private.DELETE("/users/:userID", controllers.DeleteUser(app))

	private.POST("/photos", controllers.CreatePhoto(app))
	private.GET("/photos", controllers.GetPhotos(app))
	private.PUT("/photos/:photoID", controllers.UpdatePhoto(app))
	private.DELETE("/photos/:photoID", controllers.DeletePhoto(app))

	return r
}
