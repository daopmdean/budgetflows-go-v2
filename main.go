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
	r.Use(CORSMiddleware())
	api := r.Group("/api")

	api.GET("/health", rest.HealthCheck)

	api.POST("/login", rest.Login)
	api.POST("/register", rest.Register)

	api.POST("/records/list", rest.GetUserRecords)
	api.POST("/records/recent", rest.GetRecentUserRecords)
	api.POST("/records/report", rest.ReportUserRecords)

	api.POST("/records", rest.CreateRecord)
	api.PUT("/records", rest.UpdateRecord)
	api.DELETE("/records", rest.DeleteUserRecord)

	api.POST("/records/prepare", rest.PrepareIndexes)
	api.GET("/records/partitions", rest.GetRecordPartitions)

	r.Run(":5656")
}

func setupDB(database *mongo.Database) {
	entity.IdGenDB.SetDB(database)

	entity.AppUserDB.SetDB(database)
	entity.RecordDB.SetDB(database)

	entity.RecordDBPartition.SetDBPartition(database)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers",
			"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
