package configs

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// Database Connection String (MySQL)
	DatabaseConnectionString = ""

	// HTTP Port to expose the API
	ApiHttpPort = 0

	// Key used to sign the token
	SecretKey []byte
)

// Load global parameters from environment variables
func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatalf("Error trying to load the environment variables: %v", err)
	}

	ApiHttpPort, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		log.Fatalf("Could not find the API_PORT environment variable: %v", err)
	}

	DatabaseConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_SERVER_ADDRESS"),
		os.Getenv("DATABASE_SERVER_PORT"),
		os.Getenv("DATABASE_NAME"))

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
