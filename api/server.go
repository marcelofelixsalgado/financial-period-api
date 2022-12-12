package api

import (
	"database/sql"
	"fmt"
	"log"
	"marcelofelixsalgado/financial-period-api/api/controllers/credentials"
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

	userCredentialsCreate "marcelofelixsalgado/financial-period-api/pkg/usecase/credentials/create"
	userCredentialsLogin "marcelofelixsalgado/financial-period-api/pkg/usecase/credentials/login"
	userCredentialsUpdate "marcelofelixsalgado/financial-period-api/pkg/usecase/credentials/update"

	userRepository "marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/user"
	userCreate "marcelofelixsalgado/financial-period-api/pkg/usecase/user/create"
	userDelete "marcelofelixsalgado/financial-period-api/pkg/usecase/user/delete"
	userFind "marcelofelixsalgado/financial-period-api/pkg/usecase/user/find"
	userList "marcelofelixsalgado/financial-period-api/pkg/usecase/user/list"
	userUpdate "marcelofelixsalgado/financial-period-api/pkg/usecase/user/update"

	userCredentialsRepository "marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/credentials"

	"net/http"

	"github.com/gorilla/mux"
)

func NewServer() *mux.Router {
	// Load environment variables
	configs.Load()

	// Connects to database
	databaseClient := database.NewConnection()

	userRoutes := setupUserRoutes(databaseClient)
	userCredentialsRoutes := setupUserCredentialsRoutes(databaseClient)
	periodRoutes := setupPeriodRoutes(databaseClient)
	healthRoutes := setupHealthRoutes()

	// Setup all routes
	routes := routes.NewRoutes(userCredentialsRoutes, userRoutes, periodRoutes, healthRoutes)

	router := routes.SetupRoutes()
	return router
}

func Run(router *mux.Router) {
	port := fmt.Sprintf(":%d", configs.ApiHttpPort)

	log.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func setupUserRoutes(databaseClient *sql.DB) user.UserRoutes {
	// setup respositories
	userRepository := userRepository.NewUserRepository(databaseClient)

	// setup Use Cases (services)
	createUseCase := userCreate.NewCreateUseCase(userRepository)
	deleteUseCase := userDelete.NewDeleteUseCase(userRepository)
	findUseCase := userFind.NewFindUseCase(userRepository)
	listUseCase := userList.NewListUseCase(userRepository)
	updateUseCase := userUpdate.NewUpdateUseCase(userRepository)

	// setup router handlers
	userHandler := user.NewUserHandler(createUseCase, deleteUseCase, findUseCase, listUseCase, updateUseCase)

	// setup routes
	userRoutes := user.NewUserRoutes(userHandler)

	return userRoutes
}

func setupUserCredentialsRoutes(databaseClient *sql.DB) credentials.UserCredentialsRoutes {
	// setup respository
	userRepository := userRepository.NewUserRepository(databaseClient)
	credentialsRepository := userCredentialsRepository.NewUserCredentialsRepository(databaseClient)

	// setup Use Cases (services)
	createUseCase := userCredentialsCreate.NewCreateUseCase(credentialsRepository, userRepository)
	updateUseCase := userCredentialsUpdate.NewUpdateUseCase(credentialsRepository)
	loginUseCase := userCredentialsLogin.NewLoginUseCase(credentialsRepository)

	// setup router handlers
	userCredentialsHandler := credentials.NewUserCredentialsHandler(createUseCase, updateUseCase, loginUseCase)

	// setup routes
	userCredentialsRoutes := credentials.NewUserCredentialsRoutes(userCredentialsHandler)

	return userCredentialsRoutes
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

func setupHealthRoutes() health.HealthRoutes {
	// setup router handlers
	healthHandler := health.NewHealthHandler()

	// setup routes
	healthRoutes := health.NewHealthRoutes(healthHandler)

	return healthRoutes
}
