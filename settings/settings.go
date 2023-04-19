package settings

import (
	"log"
	"os"

	"github.com/codingconcepts/env"
	"github.com/joho/godotenv"
	"github.com/marcelofelixsalgado/financial-commons/settings"
	"github.com/sirupsen/logrus"
)

// ConfigType struct to resolve env vars
type ConfigType struct {
	AppName     string `env:"APP_NAME" default:"financial-period-api"`
	Environment string `env:"ENVIRONMENT" default:"development"`

	// Database Connection String (MySQL)
	DatabaseConnectionUser          string `env:"DATABASE_USER"`
	DatabaseConnectionPassword      string `env:"DATABASE_PASSWORD"`
	DatabaseConnectionServerAddress string `env:"DATABASE_SERVER_ADDRESS"`
	DatabaseConnectionServerPort    int    `env:"DATABASE_SERVER_PORT" default:"3306"`
	DatabaseName                    string `env:"DATABASE_NAME"`

	// HTTP Port to expose the API
	ApiHttpPort int `env:"API_PORT" default:"8081"`

	// Key used to sign the token
	SecretKey []byte `env:"SECRET_KEY"`

	ServerCloseWait int `env:"SERVER_CLOSEWAIT" default:"10"`

	// Log files
	LogAccessFile string `env:"LOG_ACCESS_FILE" default:"./access.log"`
	LogAppFile    string `env:"LOG_APP_FILE" default:"./app.log"`
	LogLevel      string `env:"LOG_LEVEL" default:"INFO"`
}

var Config ConfigType

// InitConfigs initializes the environment settings
func Load() {

	settings.Load()

	// load .env (if exists)
	// err := godotenv.Load()
	loadDotEnv()

	// if err != nil {
	// 	log.Println("No .env file found")
	// }

	// bind env vars
	if err := env.Set(&Config); err != nil {
		log.Fatal(err)
	}

	if _, err := logrus.ParseLevel(Config.LogLevel); err != nil {
		log.Fatal(err)
	}
}

func loadDotEnv() {
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			logrus.Fatal("Error loading .env file")
		}
		logrus.Println("Using .env file")
	}
}
