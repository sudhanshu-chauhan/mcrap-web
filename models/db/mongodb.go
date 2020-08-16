package db

import (
	"log"
	"os"

	"github.com/goonode/mogo"
)

var mongoConnection *mogo.Connection = nil

// GetConnection is for getting mongo connection
func GetConnection() *mogo.Connection {
	if mongoConnection == nil {
		connectionString := os.Getenv("MONGO_URL")
		dbName := os.Getenv("MCRAP_WEB_MONGO_DB")
		config := &mogo.Config{
			ConnectionString: connectionString,
			Database:         dbName,
		}
		mongoConnection, err := mogo.Connect(config)
		if err != nil {
			log.Fatal(err)
		} else {
			return mongoConnection
		}
	}
	return mongoConnection
}
