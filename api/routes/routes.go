package routes

import (
	"marcelofelixsalgado/financial-period-api/api/middlewares"
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
	router.Use(middlewares.ResponseFormatMiddleware)

	// period routes
	for _, route := range periodRoutes {
		router.HandleFunc(route.URI,
			middlewares.Logger(
				middlewares.Authenticate(route.Function))).Methods(route.Method)
	}

	// health routes
	for _, route := range healthRoutes {
		router.HandleFunc(route.URI,
			middlewares.Logger(
				route.Function)).Methods(route.Method)
	}

	return router
}
