package routers

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaelzhao21/mikz.dev/util"
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
		password := util.GetEnv("PASSWORD")

		// Read data from request
		data, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Unable to read request body"})
			return
		}

		// Unmarshal JSON
		var body LoginRequest
		err = json.Unmarshal(data, &body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON request body"})
			return
		}

		// Compare passwords
		if password != body.Password {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect password"})
			return
		}

		// Generate token
		token, err = randToken()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal error while generating token"})
			return
		}

		// Return token
		ctx.JSON(http.StatusOK, gin.H{"token": token})
	})

	// Token check route; return { "ok": bool } -> true if token okay
	router.POST("/s3/token", func(ctx *gin.Context) {
		// Get token from cookies
		cookieToken, err := ctx.Cookie("token")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "token cookie not found", "ok": false})
			return
		}

		// Check token
		if cookieToken != token {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token cookie", "ok": false})
			return
		}

		// Return okay
		ctx.JSON(http.StatusOK, gin.H{"ok": true})
	})

	// Upload route
	router.POST("/s3/upload", func(ctx *gin.Context) {
		// Get file from form data
		fileHeader, err := ctx.FormFile("file")
		if err != nil {
			fmt.Println(err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid file upload"})
			return
		}

		// Open file
		file, err := fileHeader.Open()
		if err != nil {
			fmt.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to open file"})
			return
		}
		defer file.Close()

		// Read the entire file
		buffer, err := io.ReadAll(file)
		if err != nil {
			fmt.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to read file"})
			return
		}

		// Get the path from form data
		form, err := ctx.MultipartForm()
		if err != nil {
			fmt.Println(err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not read form"})
			return
		}

		pathValues := form.Value["path"]
		if len(pathValues) == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Path is missing for form"})
			return
		}
		path := pathValues[0]

		err = util.UploadToS3(buffer, path)
		if err != nil {
			fmt.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to upload to S3 bucket"})
			return
		}

		ctx.Status(http.StatusOK)
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

type LoginRequest struct {
	Password string `json:"password"`
}
