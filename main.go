package main

import "marcelofelixsalgado/financial-period-api/api"

func main() {
	// Start HTTP Server
	server := api.NewServer()
	server.Run()
}
