package api

import (
	"database/sql"
	"fmt"
	"log"
	"marcelofelixsalgado/financial-period-api/api/controllers/health"
	"marcelofelixsalgado/financial-period-api/api/controllers/period"
	"marcelofelixsalgado/financial-period-api/api/controllers/user"
	"marcelofelixsalgado/financial-period-api/api/routes"
	"marcelofelixsalgado/financial-period-api/configs"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/database"

	periodRepository "marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/period"
	periodCreate "marcelofelixsalgado/financial-period-api/pkg/usecase/period/create"
	periodDelete "marcelofelixsalgado/financial-period-api/pkg/usecase/period/delete"
	periodFind "marcelofelixsalgado/financial-period-api/pkg/usecase/period/find"
	periodList "marcelofelixsalgado/financial-period-api/pkg/usecase/period/list"
	periodUpdate "marcelofelixsalgado/financial-period-api/pkg/usecase/period/update"

	userRepository "marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/user"
	userCreate "marcelofelixsalgado/financial-period-api/pkg/usecase/user/create"
	userDelete "marcelofelixsalgado/financial-period-api/pkg/usecase/user/delete"
	userFind "marcelofelixsalgado/financial-period-api/pkg/usecase/user/find"
	userList "marcelofelixsalgado/financial-period-api/pkg/usecase/user/list"
	userUpdate "marcelofelixsalgado/financial-period-api/pkg/usecase/user/update"

	"net/http"

	"github.com/gorilla/mux"
)

func NewServer() *mux.Router {
	// Load environment variables
	configs.Load()

	// Connects to database
	databaseClient := database.NewConnection()

	periodRoutes := setupPeriodRoutes(databaseClient)
	userRoutes := setupUserRoutes(databaseClient)
	healthRoutes := setupHealthRoutes()

	// Setup all routes
	routes := routes.NewRoutes(periodRoutes, userRoutes, healthRoutes)

	router := routes.SetupRoutes()
	return router
}

func Run(router *mux.Router) {
	port := fmt.Sprintf(":%d", configs.ApiHttpPort)

	log.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func setupPeriodRoutes(databaseClient *sql.DB) period.PeriodRoutes {
	// setup respository
	repository := periodRepository.NewPeriodRepository(databaseClient)

	// setup Use Cases (services)
	createUseCase := periodCreate.NewCreateUseCase(repository)
	deleteUseCase := periodDelete.NewDeleteUseCase(repository)
	findUseCase := periodFind.NewFindUseCase(repository)
	listUseCase := periodList.NewListUseCase(repository)
	updateUseCase := periodUpdate.NewUpdateUseCase(repository)

	// setup router handlers
	periodHandler := period.NewPeriodHandler(createUseCase, deleteUseCase, findUseCase, listUseCase, updateUseCase)

	// setup routes
	periodRoutes := period.NewPeriodRoutes(periodHandler)

	return periodRoutes
}

func setupUserRoutes(databaseClient *sql.DB) user.UserRoutes {
	// setup respository
	repository := userRepository.NewUserRepository(databaseClient)

	// setup Use Cases (services)
	createUseCase := userCreate.NewCreateUseCase(repository)
	deleteUseCase := userDelete.NewDeleteUseCase(repository)
	findUseCase := userFind.NewFindUseCase(repository)
	listUseCase := userList.NewListUseCase(repository)
	updateUseCase := userUpdate.NewUpdateUseCase(repository)

	// setup router handlers
	userHandler := user.NewUserHandler(createUseCase, deleteUseCase, findUseCase, listUseCase, updateUseCase)

	// setup routes
	userRoutes := user.NewUserRoutes(userHandler)

	return userRoutes
}

func setupHealthRoutes() health.HealthRoutes {
	// setup router handlers
	healthHandler := health.NewHealthHandler()

	// setup routes
	healthRoutes := health.NewHealthRoutes(healthHandler)

	return healthRoutes
}
