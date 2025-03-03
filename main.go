package main

import (
	"github.com/gin-gonic/gin"
	"github.com/newbieDev-22/backend-challenge/controller"
	_ "github.com/newbieDev-22/backend-challenge/docs" // Import generated Swagger docs
	"github.com/newbieDev-22/backend-challenge/routes"
	"github.com/newbieDev-22/backend-challenge/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Beef Summary API
// @version         1.0
// @description     A service that provides summary statistics about beef-related text.
// @host           localhost:8080
// @BasePath       /

// @tag.name        beef
// @tag.description Operations about beef text analysis

// @schemes         http
// @produce        json

func main() {
	router := gin.Default()

	// Initialize controllers and services
	summaryBeefController := controller.NewSummaryBeefController(service.NewSummaryBeefService())

	// Setup routes
	routes.SummaryBeefRoutes(router, summaryBeefController)

	// Swagger documentation endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
