package routers

import (
	"crypto/rand"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var token string

func S3Router(router *gin.Engine) {
	// Load HTML files
	router.LoadHTMLFiles("./static/web/s3.html", "./static/web/login.html")

	// Main s3 route
	router.GET("/s3", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "s3.html", gin.H{
			"title": "S3 Upload",
		})
	})

	// Login route for s3 route
	router.GET("/s3/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{
			"title": "Login",
		})
	})

	// Login backend route
	router.POST("/s3/login", func(ctx *gin.Context) {

	})
}

// randToken generates a random token
func randToken() (string, error) {
	b := make([]byte, 12)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Printf("Error generating random number: %s\n", err.Error())
		return "", err
	}
	return fmt.Sprintf("%x", b)[2:12], nil
}
