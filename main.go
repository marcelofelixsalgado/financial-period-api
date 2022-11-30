package main

import "marcelofelixsalgado/financial-period-api/api"

func main() {
	// Start HTTP Server
	router := api.NewServer()
	api.Run(router)
}
