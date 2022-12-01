package api

import (
	"fmt"
	"log"
	"marcelofelixsalgado/financial-period-api/api/controllers/health"
	"marcelofelixsalgado/financial-period-api/api/controllers/period"
	"marcelofelixsalgado/financial-period-api/api/routes"
	"marcelofelixsalgado/financial-period-api/configs"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/database"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/period/create"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/period/delete"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/period/find"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/period/list"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/period/update"

	"net/http"

	"github.com/gorilla/mux"
)

func NewServer() *mux.Router {
	// Load environment variables
	configs.Load()

	// Connects to database
	databaseClient := database.NewConnection()

	repository := repository.NewPeriodRepository(databaseClient)

	// setup Use Cases (services)
	createUseCase := create.NewCreateUseCase(repository)
	deleteUseCase := delete.NewDeleteUseCase(repository)
	findUseCase := find.NewFindUseCase(repository)
	listUseCase := list.NewListUseCase(repository)
	updateUseCase := update.NewUpdateUseCase(repository)

	// setup router handlers
	periodHandler := period.NewPeriodHandler(createUseCase, deleteUseCase, findUseCase, listUseCase, updateUseCase)
	healthHandler := health.NewHealthHandler()

	// Setup routes
	periodRoutes := period.NewPeriodRoutes(periodHandler)
	healthRoutes := health.NewHealthRoutes(healthHandler)

	// Setup all routes
	routes := routes.NewRoutes(periodRoutes, healthRoutes)

	router := routes.SetupRoutes()
	return router
}

func Run(router *mux.Router) {
	port := fmt.Sprintf(":%d", configs.ApiHttpPort)

	log.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
