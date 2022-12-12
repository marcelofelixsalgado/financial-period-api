package routes

import (
	"marcelofelixsalgado/financial-period-api/api/controllers"
	"marcelofelixsalgado/financial-period-api/api/controllers/credentials"
	"marcelofelixsalgado/financial-period-api/api/controllers/health"
	"marcelofelixsalgado/financial-period-api/api/controllers/period"
	"marcelofelixsalgado/financial-period-api/api/controllers/user"
	"marcelofelixsalgado/financial-period-api/api/middlewares"

	"github.com/gorilla/mux"
)

type Routes struct {
	userCredentialRoutes credentials.UserCredentialsRoutes
	userRoutes           user.UserRoutes
	periodRoutes         period.PeriodRoutes
	healthRoutes         health.HealthRoutes
}

func NewRoutes(userCredentialRoutes credentials.UserCredentialsRoutes, userRoutes user.UserRoutes, periodRoutes period.PeriodRoutes, healthRoutes health.HealthRoutes) *Routes {
	return &Routes{
		userCredentialRoutes: userCredentialRoutes,
		userRoutes:           userRoutes,
		periodRoutes:         periodRoutes,
		healthRoutes:         healthRoutes,
	}
}

func (routes *Routes) SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	router.Use(middlewares.ResponseFormatMiddleware)

	// user credentials routes
	setupRoute(router, routes.userCredentialRoutes.UserCredentialsRouteMapping())
	// for _, route := range routes.userCredentialRoutes.UserCredentialsRouteMapping() {
	// 	if route.RequiresAuthentication {
	// 		router.HandleFunc(route.URI,
	// 			middlewares.Logger(
	// 				middlewares.Authenticate(route.Function))).Methods(route.Method)
	// 	} else {
	// 		router.HandleFunc(route.URI,
	// 			middlewares.Logger(route.Function)).Methods(route.Method)

	// 	}
	// }

	// user routes
	setupRoute(router, routes.userRoutes.UserRouteMapping())
	// for _, route := range routes.userRoutes.UserRouteMapping() {
	// 	router.HandleFunc(route.URI,
	// 		middlewares.Logger(
	// 			middlewares.Authenticate(route.Function))).Methods(route.Method)
	// }

	// period routes
	setupRoute(router, routes.periodRoutes.PeriodRouteMapping())
	// for _, route := range routes.periodRoutes.PeriodRouteMapping() {
	// 	router.HandleFunc(route.URI,
	// 		middlewares.Logger(
	// 			middlewares.Authenticate(route.Function))).Methods(route.Method)
	// }

	// health routes
	setupRoute(router, routes.healthRoutes.HealthRouteMapping())
	// for _, route := range routes.healthRoutes.HealthRouteMapping() {
	// 	router.HandleFunc(route.URI,
	// 		middlewares.Logger(
	// 			route.Function)).Methods(route.Method)
	// }

	return router
}

func setupRoute(router *mux.Router, routes []controllers.Route) {
	// router := mux.NewRouter()

	for _, route := range routes {
		if route.RequiresAuthentication {
			router.HandleFunc(route.URI,
				middlewares.Logger(
					middlewares.Authenticate(route.Function))).Methods(route.Method)
		} else {
			router.HandleFunc(route.URI,
				middlewares.Logger(route.Function)).Methods(route.Method)

		}
	}
	// return router
}
