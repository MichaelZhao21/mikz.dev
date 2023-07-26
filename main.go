package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/michaelzhao21/mikz.dev/routers"
)

func main() {
	// Load env
	godotenv.Load()

	// Create router instance
	router := gin.Default()

	// Cors
	router.Use(cors.Default())

	// Routers
	routers.RedirectRouter(router)
	routers.DefaultRouter(router)
	routers.S3Router(router)
	routers.DevDayRouter(router)

	// Start router
	router.Run()
}
