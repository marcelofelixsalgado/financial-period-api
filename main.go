package main

import (
	"fmt"
	"log"

	"marcelofelixsalgado/financial-month-api/api/routes"
	"net/http"

	"github.com/gorilla/mux"
)

var port = ":8081"
var basepath = "/v1/months"

func setupRoutes(router *mux.Router) {
	// POST
	// router.HandleFunc(path, routes.CreateBalances).Methods(http.MethodPost)

	// // GET
	// router.HandleFunc(path, routes.GetBalances).Methods(http.MethodGet)

	router.HandleFunc(basepath+"/health", routes.Health).Methods(http.MethodGet)
}

func main() {

	router := mux.NewRouter()

	setupRoutes(router)

	fmt.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
