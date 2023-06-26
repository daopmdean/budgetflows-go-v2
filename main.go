package main

import (
	"context"

	"github.com/daopmdean/budgetflows-go-v2/conf"
	"github.com/daopmdean/budgetflows-go-v2/entity"
	"github.com/daopmdean/budgetflows-go-v2/rest"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	conf.InitAppConfig()

	mongoClient, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(conf.AppConfig.MongoUri),
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
	api := r.Group("/api")

	api.GET("/health", rest.HealthCheck)

	api.POST("/login", rest.Login)
	api.POST("/register", rest.Register)

	api.POST("/records", rest.CreateRecord)
	api.POST("/records/list", rest.GetUserRecords)
	api.POST("/records/report", rest.ReportUserRecords)
	api.DELETE("/records", rest.DeleteUserRecord)
	api.POST("/records/prepare", rest.PrepareIndexes)

	r.Run(":5656")
}

func setupDB(database *mongo.Database) {
	entity.IdGenDB.SetDB(database)

	entity.AppUserDB.SetDB(database)
	entity.RecordDB.SetDB(database)

	entity.RecordDBPartition.SetDBPartition(database)
}
