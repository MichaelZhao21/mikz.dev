package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultRouter(router *gin.Engine) {
	// Default endpoint
	router.StaticFile("/", "./static/aaaaa.webp")

	// Load HTML
	router.LoadHTMLFiles("./static/web/dance.html")

	// Serve static files
	router.Static("/static", "./static/web")

	// Dance endpoint
	router.GET("/dance", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "dance.html", gin.H{
			"title": "Dancing!",
		})
	})

	// Ping pong endpoint
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

}
