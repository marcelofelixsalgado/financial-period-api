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
)

// Load global parameters from environment variables
func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	ApiHttpPort, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	DatabaseConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_SERVER_ADDRESS"),
		os.Getenv("DATABASE_SERVER_PORT"),
		os.Getenv("DATABASE_NAME"))
}
