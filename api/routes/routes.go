package routes

import (
	"marcelofelixsalgado/financial-period-api/api/controllers/health"
	"marcelofelixsalgado/financial-period-api/api/controllers/login"
	"marcelofelixsalgado/financial-period-api/api/controllers/period"
	"marcelofelixsalgado/financial-period-api/api/controllers/user"
	"marcelofelixsalgado/financial-period-api/api/middlewares"

	"github.com/gorilla/mux"
)

type Routes struct {
	loginRoutes  login.LoginRoutes
	userRoutes   user.UserRoutes
	periodRoutes period.PeriodRoutes
	healthRoutes health.HealthRoutes
}

func NewRoutes(loginRoutes login.LoginRoutes, userRoutes user.UserRoutes, periodRoutes period.PeriodRoutes, healthRoutes health.HealthRoutes) *Routes {
	return &Routes{
		loginRoutes:  loginRoutes,
		userRoutes:   userRoutes,
		periodRoutes: periodRoutes,
		healthRoutes: healthRoutes,
	}
}

func (routes *Routes) SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	router.Use(middlewares.ResponseFormatMiddleware)

	// login routes
	for _, route := range routes.loginRoutes.LoginRouteMapping() {
		router.HandleFunc(route.URI,
			middlewares.Logger(
				middlewares.Authenticate(route.Function))).Methods(route.Method)
	}

	// user routes
	for _, route := range routes.userRoutes.UserRouteMapping() {
		router.HandleFunc(route.URI,
			middlewares.Logger(
				middlewares.Authenticate(route.Function))).Methods(route.Method)
	}

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
