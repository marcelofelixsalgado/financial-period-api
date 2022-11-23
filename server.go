package main

import (
	"fmt"
	"log"
	"marcelofelixsalgado/financial-period-api/api/routes"
	"marcelofelixsalgado/financial-period-api/configs"

	"net/http"
)

func startServer() {
	port := fmt.Sprintf(":%d", configs.ApiHttpPort)

	router := routes.SetupRoutes()

	fmt.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
