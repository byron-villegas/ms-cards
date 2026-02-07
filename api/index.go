package main

import (
	"ms-cards/config"
	"ms-cards/routes"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	// Show banner with application information
	config.ShowBanner()

	// Create a server instance
	router = gin.Default()

	// Config CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.AppConfig.CORSOrigin},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{config.AppConfig.CORSHeaders},
		AllowCredentials: true,
	}))

	// Config server path
	routerGroup := router.Group(config.AppConfig.ServerPath)

	// Config routes for the server
	routes.SetupRoutes(routerGroup)
}

// Handler is the entry point for Vercel serverless function
func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}