package database

import (
	"errors"
	"fmt"
	"os"

	"github.com/prem/callcenter/sharedlib/repository"
	"go.uber.org/zap"
)

func GetMongoConnection(collectionName string) (repository.MongodbConnection, error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	var mongoDbUri string
	debugMode := os.Getenv("DEBUG_MODE")

	if len(debugMode) == 0 || debugMode == "false" {
		mongoDbUri = os.Getenv("CC_MONGO_SERVER")
		if len(mongoDbUri) == 0 {
			fmt.Println("CC_MONGO_SERVER is missing.")
			logger.Warn("CC_MONGO_SERVER is missing.")
			return repository.MongodbConnection{}, errors.New("CC_MONGO_SERVER is missing")
		}
	} else {
		mongoDbUri = os.Getenv("MONGO_SERVER_URI")
		if len(mongoDbUri) == 0 {
			fmt.Println("MONGO_SERVER_URI is missing.")
			return repository.MongodbConnection{}, errors.New("MONGO_SERVER_URI is missing.")
		}
	}

	fmt.Println("mongodb-URI:" + mongoDbUri)
	logger.Info("mongodb-URI:" + mongoDbUri)
	return repository.MongodbConnection{Uri: mongoDbUri, DbName: "callcenter", CollectionName: collectionName}, nil
}
