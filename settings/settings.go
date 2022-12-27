package settings

import (
	"log"

	"github.com/codingconcepts/env"
	"github.com/joho/godotenv"
)

// var (
// 	// Database Connection String (MySQL)
// 	DatabaseConnectionString = ""

// 	// HTTP Port to expose the API
// 	ApiHttpPort = 0

// 	// Key used to sign the token
// 	SecretKey []byte
// )

// // Load global parameters from environment variables
// func Load() {
// 	var err error

// 	if err = godotenv.Load(); err != nil {
// 		log.Fatalf("Error trying to load the environment variables: %v", err)
// 	}

// 	ApiHttpPort, err = strconv.Atoi(os.Getenv("API_PORT"))
// 	if err != nil {
// 		log.Fatalf("Could not find the API_PORT environment variable: %v", err)
// 	}

// 	DatabaseConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
// 		os.Getenv("DATABASE_USER"),
// 		os.Getenv("DATABASE_PASSWORD"),
// 		os.Getenv("DATABASE_SERVER_ADDRESS"),
// 		os.Getenv("DATABASE_SERVER_PORT"),
// 		os.Getenv("DATABASE_NAME"))

// 	SecretKey = []byte(os.Getenv("SECRET_KEY"))
// }

// ConfigType struct to resolve env vars
type ConfigType struct {
	AppName     string `env:"APP_NAME" default:"financial-period-api"`
	Environment string `env:"ENVIRONMENT" default:"development"`

	// Database Connection String (MySQL)
	DatabaseConnectionUser          string `env:"DATABASE_USER"`
	DatabaseConnectionPassword      string `env:"DATABASE_PASSWORD"`
	DatabaseConnectionServerAddress string `env:"DATABASE_SERVER_ADDRESS"`
	DatabaseConnectionServerPort    string `env:"DATABASE_SERVER_PORT" default:"3306"`
	DatabaseName                    string `env:"DATABASE_NAME"`

	// HTTP Port to expose the API
	ApiHttpPort int `env:"API_PORT" default:"8081"`

	// Key used to sign the token
	SecretKey []byte `env:"SECRET_KEY"`

	ServerCloseWait int `env:"SERVER_CLOSEWAIT" default:"10"`
}

var Config ConfigType

// InitConfigs initializes the environment settings
func Load() {
	// load .env (if exists)
	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found")
	}

	// bind env vars
	if err := env.Set(&Config); err != nil {
		log.Fatal(err)
	}
}
