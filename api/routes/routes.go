package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI                    string
	Method                 string
	Function               func(w http.ResponseWriter, r *http.Request)
	RequiresAuthentication bool
}

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	router.Use(responseFormatMiddleware)

	// period routes
	for _, route := range periodRoutes {
		router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	// health routes
	for _, route := range healthRoutes {
		router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return router
}

func responseFormatMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
