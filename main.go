package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/daopmdean/budgetflows-go-v2/conf"
	"github.com/daopmdean/budgetflows-go-v2/entity"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	serverStartTime := time.Now()

	appConfig := conf.Config{
		MongoUri:  os.Getenv("MONGO_URI"),
		SignedKey: os.Getenv("SECRET_KEY"),
	}

	mongoClient, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(appConfig.MongoUri),
	)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	database := mongoClient.Database("budgetflows")

	setupDB(database)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":   "I'm alive!",
			"startTime": serverStartTime,
			"timeNow":   time.Now(),
		})
	})

	r.Run(":5656")
}

func setupDB(database *mongo.Database) {
	entity.AppUserDB.SetDB(database)
	entity.RecordDB.SetDB(database)
}
