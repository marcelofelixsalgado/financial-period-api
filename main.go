package main

import (
	"fmt"
	"log"

	"marcelofelixsalgado/financial-period-api/api/routes"
	"net/http"

	"github.com/gorilla/mux"
)

var port = ":8081"
var basepath = "/v1/periods"

func setupRoutes(router *mux.Router) {
	// POST
	router.HandleFunc(basepath, routes.CreatePeriod).Methods(http.MethodPost)

	// GET
	router.HandleFunc(basepath, routes.ListPeriods).Methods(http.MethodGet)

	// GET
	router.HandleFunc(basepath+"/{id}", routes.GetPeriodById).Methods(http.MethodGet)

	// HEALTH
	router.HandleFunc(basepath+"/health", routes.Health).Methods(http.MethodGet)
}

func responseFormatMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {

	router := mux.NewRouter()
	router.Use(responseFormatMiddleware)

	setupRoutes(router)

	fmt.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
