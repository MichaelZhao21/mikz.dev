package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/michaelzhao21/mikz.dev/routers"
)

func main() {
	router := gin.Default()

	// Cors
	router.Use(cors.Default())

	// Routers
	routers.RedirectRouter(router)
	routers.DefaultRouter(router)

	router.Run()
}
