package main

import (
	"ms-cards/config"
	"ms-cards/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "ms-cards/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           MS Cards API
// @version         1.0
// @description     Microservice for managing trading cards
// @host            localhost:8080
// @BasePath        /api
// @schemes         http https
func main() {
	// Show banner with application information
	config.ShowBanner()

	// Create a server instance
	r := gin.Default()

	// Config CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.AppConfig.CORSOrigin},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{config.AppConfig.CORSHeaders},
		AllowCredentials: true,
	}))

	// Config server path
	routerGroup := r.Group(config.AppConfig.ServerPath)

	// Config routes for the server
	routes.SetupRoutes(routerGroup)

	// Swagger documentation route
	routerGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server with port
	r.Run(":" + config.AppConfig.ServerPort)
}
