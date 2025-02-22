package handlers

import (
	"fmt"
	"os"
	"teams/models"

	"github.com/gin-gonic/gin"
)

func HealthHandler(c *gin.Context) {
	fmt.Println("Env-FileName:", os.Getenv("FILEPATH")+os.Getenv("FILENAME"))

	var apiSettings struct {
		FilePath   string
		FileName   string
		MongoPort  string
		MongoURI   string
		Version    string
		LastUpdate string
	}

	apiSettings.FileName = os.Getenv("FILENAME")
	apiSettings.FilePath = os.Getenv("FILEPATH")
	apiSettings.Version = os.Getenv("VERSION")
	apiSettings.MongoPort = os.Getenv("MONGO_SVC_PORT")
	apiSettings.MongoURI = os.Getenv("MONGO_SERVER_URI")
	apiSettings.LastUpdate = os.Getenv("LASTUPDATE")

	response := models.ApiResponse{
		Status:  "200",
		Message: "Web-app is healthy...",
		Data: []interface{}{
			apiSettings,
		},
	}

	c.JSON(200, response)
}
