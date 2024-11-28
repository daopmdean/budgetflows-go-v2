package conf

import (
	"os"
	"time"
)

var AppConfig *Config

type Config struct {
	AppName         string
	Port            int
	Env             string
	MongoUri        string
	SignedKey       string
	ServerStartTime time.Time
}

func InitAppConfig() {

	appConfig := Config{
		AppName:         "budgetflows-go-v2",
		Port:            5656,
		Env:             os.Getenv("ENV"), // dev, prod
		MongoUri:        os.Getenv("MONGO_URI"),
		SignedKey:       os.Getenv("SECRET_KEY"),
		ServerStartTime: time.Now(),
	}

	AppConfig = &appConfig
}
