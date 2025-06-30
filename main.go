package main

import (
	"gotest/core"
	"gotest/models"
	"gotest/routes"
)

// @title My Gin API
// @version 1.0
// @description Example API using Gin, JWT, GORM, Swagger

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	core.InitDB()
	core.DB.AutoMigrate(&models.User{})
	r := routes.SetupRouter()
	r.Run(":3000")
}
