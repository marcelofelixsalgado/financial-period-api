package main

import (
	"fmt"
	"log"
	"marcelofelixsalgado/financial-period-api/api/routes"

	"net/http"
)

var port = ":8081"

func startServer() {

	router := routes.SetupRoutes()

	fmt.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
