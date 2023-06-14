package conf

import (
	"os"
	"time"
)

var AppConfig Config

type Config struct {
	MongoUri        string
	SignedKey       string
	ServerStartTime time.Time
}

func InitAppConfig() {

	appConfig := Config{
		MongoUri:        os.Getenv("MONGO_URI"),
		SignedKey:       os.Getenv("SECRET_KEY"),
		ServerStartTime: time.Now(),
	}

	AppConfig = appConfig
}
