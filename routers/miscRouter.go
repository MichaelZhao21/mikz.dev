package routers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	CDN_URL = "https://cdn.michaelzhao.xyz/hackutd-devday-workshop/"
)

func MiscRouter(router *gin.Engine) {
	// GET /devday/help - returns the API reference
	router.GET("/devday/help", func(ctx *gin.Context) {
		videoFormat := gin.H{
			"id":           "string",
			"title":        "string",
			"thumbnailUrl": "string",
			"videoUrl":     "string",
		}
		body := gin.H{
			"GET /": gin.H{
				"response": []gin.H{videoFormat},
			},
			"GET /{id}": gin.H{
				"response": videoFormat,
			},
		}
		ctx.JSON(http.StatusOK, body)
	})

	// GET /devday - returns a list of videos
	router.GET("/devday", func(ctx *gin.Context) {
		// Read JSON file
		data, err := os.ReadFile("./static/devDay.json")
		if err != nil {
			log.Printf("Invalid devDay.json file\n")
			ctx.Status(http.StatusInternalServerError)
			return
		}

		// Unmarshal the JSON
		var videos []Video
		err = json.Unmarshal(data, &videos)
		if err != nil {
			log.Printf("Invalid devDay.json file\n")
			ctx.Status(http.StatusInternalServerError)
			return
		}

		// Update the video URLs
		for idx := range videos {
			updateUrls(&videos[idx])
		}

		// Send data to user
		ctx.JSON(http.StatusOK, videos)
	})

	// GET /devday/{id} - returns a single video given its ID
	router.GET("/devday/:id", func(ctx *gin.Context) {
		// Read JSON file
		data, err := os.ReadFile("./static/devDay.json")
		if err != nil {
			log.Printf("Invalid devDay.json file\n")
			ctx.Status(http.StatusInternalServerError)
			return
		}

		// Unmarshal the JSON
		var videos []Video
		err = json.Unmarshal(data, &videos)
		if err != nil {
			log.Printf("Invalid devDay.json file\n")
			ctx.Status(http.StatusInternalServerError)
			return
		}

		// Get the ID
		id := ctx.Param("id")

		// Search through the video file and return the matching ID
		for _, video := range videos {
			if video.Id == id {
				updateUrls(&video)
				ctx.JSON(http.StatusOK, video)
				return
			}
		}

		// If no video found, return 400
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Video with ID %s not found", id),
		})
	})
}

type Video struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	ThumbnailUrl string `json:"thumbnailUrl"`
	VideoUrl     string `json:"videoUrl"`
}

// updateUrls will combine the CDN_URL with the given
// thumbnail and video urls
func updateUrls(video *Video) {
	video.ThumbnailUrl = CDN_URL + video.ThumbnailUrl
	video.VideoUrl = CDN_URL + video.VideoUrl
}
