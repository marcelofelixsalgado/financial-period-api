package routes

import (
	"marcelofelixsalgado/financial-period-api/api/controllers/health"
	"marcelofelixsalgado/financial-period-api/api/controllers/period"
	"marcelofelixsalgado/financial-period-api/api/middlewares"

	"github.com/gorilla/mux"
)

type Routes struct {
	periodRoutes period.PeriodRoutes
	healthRoutes health.HealthRoutes
}

func NewRoutes(periodRoutes period.PeriodRoutes, healthRoutes health.HealthRoutes) *Routes {
	return &Routes{
		periodRoutes: periodRoutes,
		healthRoutes: healthRoutes,
	}
}

func (routes *Routes) SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	router.Use(middlewares.ResponseFormatMiddleware)

	// period routes
	for _, route := range routes.periodRoutes.PeriodRouteMapping() {
		router.HandleFunc(route.URI,
			middlewares.Logger(
				middlewares.Authenticate(route.Function))).Methods(route.Method)
	}

	// health routes
	for _, route := range routes.healthRoutes.HealthRouteMapping() {
		router.HandleFunc(route.URI,
			middlewares.Logger(
				route.Function)).Methods(route.Method)
	}

	return router
}
