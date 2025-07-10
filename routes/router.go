package routes

import (
	"gotest/controllers"
	_ "gotest/docs"
	"gotest/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/login", controllers.Login)

	auth := r.Group("/")
	auth.Use(middleware.JWTAuth())
	auth.GET("/protected", controllers.Protected)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// User routes
	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.ListUsers)
	r.DELETE("/users/:id", controllers.DeleteUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.GET("/users/search", controllers.SearchUsers)
	r.GET("/users/count", controllers.CountUsers)
	r.GET("/users/email", controllers.GetUserByEmail)
	r.GET("/users/name", controllers.GetUserByUsername)
	r.GET("/users/:id", controllers.GetUserByID)
	r.GET("/users/transfer", controllers.Transfer)

	r.GET("/test/", controllers.GetTest)

	return r
}
