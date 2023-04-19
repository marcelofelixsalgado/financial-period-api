package main

import "github.com/marcelofelixsalgado/financial-period-api/api"

func main() {
	// Start HTTP Server
	server := api.NewServer()
	server.Run()
}
