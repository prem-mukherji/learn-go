package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"teams/database"
	"teams/viewModels"

	"github.com/gin-gonic/gin"
	"github.com/prem/callcenter/sharedlib/dbentities"
	"go.uber.org/zap"
)

const APPPrefix = "ccTeams"

func GetTeamsFromFileHandler(c *gin.Context) {
	// Open jsonFile
	jsonFile, err := os.Open(os.Getenv("FILEPATH") + os.Getenv("FILENAME"))
	// if os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var result map[string]interface{}

	json.Unmarshal([]byte(byteValue), &result)

	array := result["teams"].([]interface{})

	fmt.Println(array)

	for i := range array {
		fmt.Println(array[i])
		//fmt.Printf("key=%s value=%+v\n ", k, v)
	}

	response := viewModels.ApiResponse{
		Status:  os.Getenv("FILEPATH") + os.Getenv("FILENAME"),
		Message: "Data retrieved successfully",
		Data:    array,
	}

	c.JSON(200, response)
}

func GetTeamsHandler_1(c *gin.Context) {
	var userlist []interface{}

	userlist = append(userlist,
		viewModels.UserData{
			Username: "prem",
			Email:    "prem@example.com",
		})

	userlist = append(userlist,
		viewModels.UserData{
			Username: "priya",
			Email:    "priya@example.com",
		})

	response := viewModels.ApiResponse{
		Status:  os.Getenv("NEWVAR"),
		Message: "Data retrieved successfully",
		Data:    userlist,
	}

	c.JSON(200, response)
}

func GetTeamsHandler(c *gin.Context) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	response := viewModels.ApiResponse{
		Status:  "400",
		Message: APPPrefix + "Unable to retrieve data.",
	}

	mongoConn, err := database.GetMongoConnection("team")
	if err != nil {
		response.Message += "mongo connection error: " + err.Error()
		logger.Error(response.Message)
		c.JSON(400, response)
		return
	}

	var teams []dbentities.Team
	filter := dbentities.Team{}

	//--------------Get the existing secret entity--------------
	findErr := mongoConn.Find(filter, &teams, nil)
	if findErr != nil {
		if findErr.Error() == "mongo: no documents in result" {
			response.Message += "Teams not found in DB"
			logger.Error(response.Message)
			c.JSON(400, response)
			return
		} else {
			response.Message += "error in reading data: " + findErr.Error()
			logger.Error(response.Message)
			c.JSON(400, response)
			return
		}
	}

	var array []interface{}

	for _, v := range teams {
		array = append(array, v)
	}

	response = viewModels.ApiResponse{
		Status:  "200",
		Message: "Data retrieved successfully",
		Data:    array,
	}

	c.JSON(200, response)
}
