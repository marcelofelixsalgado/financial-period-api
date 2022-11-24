package main

import (
	"marcelofelixsalgado/financial-period-api/configs"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/database"
)

func main() {
	// Load environment variables
	configs.Load()

	// Connects to database
	database.Connect()

	// Start HTTP Server
	startServer()
}
