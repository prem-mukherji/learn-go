package main

import (
	"teams/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	SetupServer().Run()
}

func SetupServer() *gin.Engine {

	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.GET("/", handlers.HealthHandler)
	r.GET("/teamsFromFile", handlers.GetTeamsFromFileHandler)
	r.GET("/teams", handlers.GetTeamsHandler)
	r.POST("/teams", handlers.CreateTeamHandler)
	return r
}
