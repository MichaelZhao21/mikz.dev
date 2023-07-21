package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultRouter(router *gin.Engine) {
	// Default endpoint
	router.StaticFile("/", "./static/aaaaa.webp")

	// Ping pong endpoint
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

}
