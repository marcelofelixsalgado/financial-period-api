package main

import (
	"marcelofelixsalgado/financial-period-api/configs"
)

func main() {
	// Load environment variables
	configs.Load()

	// Start HTTP Server
	startServer()
}
