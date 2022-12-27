package api

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"marcelofelixsalgado/financial-period-api/api/controllers/health"
	"marcelofelixsalgado/financial-period-api/api/controllers/login"
	"marcelofelixsalgado/financial-period-api/api/controllers/period"
	"marcelofelixsalgado/financial-period-api/api/controllers/user"
	"marcelofelixsalgado/financial-period-api/api/routes"
	"marcelofelixsalgado/financial-period-api/commons/logger"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/database"
	"marcelofelixsalgado/financial-period-api/settings"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	logs "marcelofelixsalgado/financial-period-api/commons/logger"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

// Server this is responsible for running an http server
type Server struct {
	http   *echo.Echo
	routes *routes.Routes
	stop   chan struct{}
}

// Start - Entry point of the API
func NewServer() *Server {
	// Load environment variables
	settings.Load()

	server := &Server{
		stop: make(chan struct{}),
	}

	return server
}

// Run is the procedure main for start the application
func (s *Server) Run() {
	s.startServer()
	<-s.stop
}

func (server *Server) startServer() {
	go server.watchStop()

	server.http = echo.New()
	logger := logs.GetLogger()
	logger.Infof("Server is starting now in %s.", settings.Config.Environment)

	// Middlewares
	server.http.Use(echoMiddleware.Logger())

	// Connects to database
	databaseClient := database.NewConnection()

	userRoutes := setupUserRoutes(databaseClient)
	loginRoutes := setupLoginRoutes(databaseClient)
	periodRoutes := setupPeriodRoutes(databaseClient)
	healthRoutes := setupHealthRoutes()

	// // Setup all routes
	routes := routes.NewRoutes(loginRoutes, userRoutes, periodRoutes, healthRoutes)

	routes.RouteMapping(server.http)
	server.routes = routes

	showRoutes(server.http)

	addr := fmt.Sprintf(":%v", settings.Config.ApiHttpPort)
	go func() {
		if err := server.http.Start(addr); err != nil {
			log.Printf("Shutting down the server now")
		}
	}()
}

// watchStop wait for the interrupt signal.
func (server *Server) watchStop() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-stop)
	server.stopServer()
}

// stopServer stops the server http
func (s *Server) stopServer() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(settings.Config.ServerCloseWait))
	defer cancel()

	logger := logs.GetLogger()
	logger.Info("Server is stoping...")
	s.http.Shutdown(ctx)
	close(s.stop)
}

func setupUserRoutes(databaseClient *sql.DB) user.UserRoutes {
	// setup respositories
	userRepository := userRepository.NewUserRepository(databaseClient)
	credentialsRepository := userCredentialsRepository.NewUserCredentialsRepository(databaseClient)

	// setup Use Cases (services)
	userCreateUseCase := userCreate.NewCreateUseCase(userRepository)
	userDeleteUseCase := userDelete.NewDeleteUseCase(userRepository)
	userFindUseCase := userFind.NewFindUseCase(userRepository)
	userListUseCase := userList.NewListUseCase(userRepository)
	userUpdateUseCase := userUpdate.NewUpdateUseCase(userRepository)

	userCredentialsCreateUseCase := userCredentialsCreate.NewCreateUseCase(credentialsRepository, userRepository)
	userCredentialsUpdateUseCase := userCredentialsUpdate.NewUpdateUseCase(credentialsRepository)

	// setup router handlers
	userHandler := user.NewUserHandler(userCreateUseCase, userDeleteUseCase, userFindUseCase, userListUseCase, userUpdateUseCase,
		userCredentialsCreateUseCase, userCredentialsUpdateUseCase)

	// setup routes
	userRoutes := user.NewUserRoutes(userHandler)

	return userRoutes
}

func setupLoginRoutes(databaseClient *sql.DB) login.LoginRoutes {
	// setup respository
	credentialsRepository := userCredentialsRepository.NewUserCredentialsRepository(databaseClient)

	// setup Use Cases (services)
	loginUseCase := userCredentialsLogin.NewLoginUseCase(credentialsRepository)

	// setup router handlers
	loginHandler := login.NewLoginHandler(loginUseCase)

	// setup routes
	loginRoutes := login.NewLoginRoutes(loginHandler)

	return loginRoutes
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

func showRoutes(e *echo.Echo) {
	var routes = e.Routes()
	logger := logger.GetLogger()

	if len(routes) > 0 {
		for _, route := range routes {
			// if strings.Contains(route.Name, "forklift-api") {
			logger.Infof("%6s: %s \n", route.Method, route.Path)
			// }
		}
	}
}
