package routers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type RedirectRoute struct {
	Slug string `json:"slug"` // Should start with /
	Url  string `json:"url"`  // Full URL to redirect to
}

func RedirectRouter(router *gin.Engine) {
	// Read in the JSON file
	dat, err := os.ReadFile("./static/redirectRoutes.json")
	if err != nil {
		log.Fatalf("Error reading routes: %s\n", err.Error())
	}

	// Convert to JSON struct
	var redirectRoutes []RedirectRoute
	err = json.Unmarshal(dat, &redirectRoutes)
	if err != nil {
		log.Fatalf("Error parsing routes to JSON: %s\n", err.Error())
	}

	// Build redirect routes
	for _, route := range redirectRoutes {
		router.GET(route.Slug, func(ctx *gin.Context) {
			ctx.Redirect(http.StatusMultipleChoices, route.Url)
		})
	}
}
