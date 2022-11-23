package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

var basepath = "/v1/periods"

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	router.Use(responseFormatMiddleware)

	// POST
	router.HandleFunc(basepath, createPeriod).Methods(http.MethodPost)

	// GET
	router.HandleFunc(basepath, listPeriods).Methods(http.MethodGet)

	// GET
	router.HandleFunc(basepath+"/{id}", getPeriodById).Methods(http.MethodGet)

	// PUT
	router.HandleFunc(basepath+"/{id}", updatePeriod).Methods(http.MethodPut)

	// DELETE
	router.HandleFunc(basepath+"/{id}", deletePeriod).Methods(http.MethodDelete)

	// HEALTH
	router.HandleFunc(basepath+"/health", health).Methods(http.MethodGet)

	return router
}

func responseFormatMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
