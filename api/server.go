package api

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/marcelofelixsalgado/financial-period-api/api/controllers/balance"

	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/marcelofelixsalgado/financial-period-api/api/controllers/category"
	"github.com/marcelofelixsalgado/financial-period-api/api/controllers/health"
	"github.com/marcelofelixsalgado/financial-period-api/api/controllers/login"
	"github.com/marcelofelixsalgado/financial-period-api/api/controllers/period"
	"github.com/marcelofelixsalgado/financial-period-api/api/controllers/subcategory"
	"github.com/marcelofelixsalgado/financial-period-api/api/controllers/transactiontype"
	"github.com/marcelofelixsalgado/financial-period-api/api/controllers/user"
	"github.com/marcelofelixsalgado/financial-period-api/api/middlewares"
	"github.com/marcelofelixsalgado/financial-period-api/api/routes"
	"github.com/marcelofelixsalgado/financial-period-api/commons/logger"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/database"
	"github.com/marcelofelixsalgado/financial-period-api/settings"

	periodRepository "github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/period"
	periodCreate "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/period/create"
	periodDelete "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/period/delete"
	periodFind "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/period/findbyid"
	periodList "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/period/list"
	periodUpdate "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/period/update"

	userCredentialsCreate "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/credentials/create"
	userCredentialsLogin "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/credentials/login"
	userCredentialsUpdate "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/credentials/update"

	tenantRepository "github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/tenant"

	transactionTypeRepository "github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/transactiontype"
	transactionTypeFind "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/transactiontype/find"
	transactionTypeList "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/transactiontype/list"

	categoryRepository "github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/category"
	categoryCreate "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/category/create"
	categoryDelete "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/category/delete"
	categoryFind "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/category/find"
	categoryList "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/category/list"
	categoryUpdate "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/category/update"

	subCategoryRepository "github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/subcategory"
	subCategoryCreate "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/subcategory/create"
	subCategoryDelete "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/subcategory/delete"
	subCategoryFind "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/subcategory/find"
	subCategoryList "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/subcategory/list"
	subCategoryUpdate "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/subcategory/update"

	userRepository "github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/user"
	userCreate "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/user/create"
	userDelete "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/user/delete"
	userFind "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/user/find"
	userList "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/user/list"
	userUpdate "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/user/update"

	balanceRepository "github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/balance"
	balanceCreate "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/balance/create"
	balanceDelete "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/balance/delete"
	balanceFind "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/balance/find"
	balanceList "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/balance/list"
	balanceUpdate "github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/balance/update"

	userCredentialsRepository "github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/credentials"

	"github.com/labstack/echo/v4"
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
	logger := logger.GetLogger()

	logger.Infof("Server is starting now in %s.", settings.Config.Environment)

	// Middlewares
	server.http.Use(middlewares.Logger())

	fmt.Println(settings.Config.DatabaseConnectionUser)
	fmt.Println(settings.Config.DatabaseConnectionPassword)
	fmt.Println(settings.Config.DatabaseConnectionServerAddress)
	fmt.Println(settings.Config.DatabaseConnectionServerPort)
	fmt.Println(settings.Config.DatabaseName)

	// Connects to database
	databaseClient := database.NewConnection()

	userRoutes := setupUserRoutes(databaseClient)
	loginRoutes := setupLoginRoutes(databaseClient)
	transactionTypeRoutes := setupTransactionTypeRoutes(databaseClient)
	categoryRoutes := setupCategoryRoutes(databaseClient)
	subCategoryRoutes := setupSubCategoryRoutes(databaseClient)
	periodRoutes := setupPeriodRoutes(databaseClient)
	balanceRoutes := setupBalanceRoutes(databaseClient)
	healthRoutes := setupHealthRoutes()

	// Setup all routes
	routes := routes.NewRoutes(loginRoutes, userRoutes, transactionTypeRoutes, categoryRoutes, subCategoryRoutes, periodRoutes, balanceRoutes, healthRoutes)

	routes.RouteMapping(server.http)
	server.routes = routes

	showRoutes(server.http)

	addr := fmt.Sprintf(":%v", settings.Config.ApiHttpPort)

	go func() {
		if err := server.http.Start(addr); err != nil {
			logger.Errorf("Shutting down the server now: %s", err)
		}
	}()
}

// watchStop wait for the interrupt signal.
func (server *Server) watchStop() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	logger.GetLogger().Info(<-stop)
	server.stopServer()
}

// stopServer stops the server http
func (s *Server) stopServer() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(settings.Config.ServerCloseWait))
	defer cancel()

	logger := logger.GetLogger()
	logger.Info("Server is stoping...")
	s.http.Shutdown(ctx)
	close(s.stop)
}

func setupUserRoutes(databaseClient *sql.DB) user.UserRoutes {
	// setup respositories
	tenantRepository := tenantRepository.NewTenantRepository(databaseClient)
	userRepository := userRepository.NewUserRepository(databaseClient)
	credentialsRepository := userCredentialsRepository.NewUserCredentialsRepository(databaseClient)

	// setup Use Cases (services)
	userCreateUseCase := userCreate.NewCreateUseCase(userRepository, tenantRepository)
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

func setupTransactionTypeRoutes(databaseClient *sql.DB) transactiontype.TransactionTypeRoutes {
	// setup respository
	repository := transactionTypeRepository.NewTransactionTypeRepository(databaseClient)

	// setup Use Cases (services)
	findUseCase := transactionTypeFind.NewFindUseCase(repository)
	listUseCase := transactionTypeList.NewListUseCase(repository)

	// setup router handlers
	handler := transactiontype.NewTransactionTypeHandler(findUseCase, listUseCase)

	// setup routes
	routes := transactiontype.NewTransactionTypeRoutes(handler)

	return routes
}

func setupCategoryRoutes(databaseClient *sql.DB) category.CategoryRoutes {
	// setup respositories
	categoryRepository := categoryRepository.NewCategoryRepository(databaseClient)
	transactionTypeRepository := transactionTypeRepository.NewTransactionTypeRepository(databaseClient)

	// setup Use Cases (services)
	createUseCase := categoryCreate.NewCreateUseCase(categoryRepository, transactionTypeRepository)
	deleteUseCase := categoryDelete.NewDeleteUseCase(categoryRepository)
	findUseCase := categoryFind.NewFindUseCase(categoryRepository)
	listUseCase := categoryList.NewListUseCase(categoryRepository)
	updateUseCase := categoryUpdate.NewUpdateUseCase(categoryRepository)

	// setup router handlers
	handler := category.NewCategoryHandler(createUseCase, deleteUseCase, findUseCase, listUseCase, updateUseCase)

	// setup routes
	routes := category.NewCategoryRoutes(handler)

	return routes
}

func setupSubCategoryRoutes(databaseClient *sql.DB) subcategory.SubCategoryRoutes {
	// setup respositories
	categoryRepository := categoryRepository.NewCategoryRepository(databaseClient)
	subCategoryRepository := subCategoryRepository.NewSubCategoryRepository(databaseClient)

	// setup Use Cases (services)
	createUseCase := subCategoryCreate.NewCreateUseCase(subCategoryRepository, categoryRepository)
	deleteUseCase := subCategoryDelete.NewDeleteUseCase(subCategoryRepository)
	findUseCase := subCategoryFind.NewFindUseCase(subCategoryRepository)
	listUseCase := subCategoryList.NewListUseCase(subCategoryRepository)
	updateUseCase := subCategoryUpdate.NewUpdateUseCase(subCategoryRepository, categoryRepository)

	// setup router handlers
	handler := subcategory.NewSubCategoryHandler(createUseCase, deleteUseCase, findUseCase, listUseCase, updateUseCase)

	// setup routes
	routes := subcategory.NewSubCategoryRoutes(handler)

	return routes
}

func setupPeriodRoutes(databaseClient *sql.DB) period.PeriodRoutes {
	// setup respository
	repository := periodRepository.NewPeriodRepository(databaseClient)

	// setup Use Cases (services)
	createUseCase := periodCreate.NewCreateUseCase(repository)
	deleteUseCase := periodDelete.NewDeleteUseCase(repository)
	findByIdUseCase := periodFind.NewFindByIdUseCase(repository)
	listUseCase := periodList.NewListUseCase(repository)
	updateUseCase := periodUpdate.NewUpdateUseCase(repository)

	// setup router handlers
	periodHandler := period.NewPeriodHandler(createUseCase, deleteUseCase, findByIdUseCase, listUseCase, updateUseCase)

	// setup routes
	periodRoutes := period.NewPeriodRoutes(periodHandler)

	return periodRoutes
}

func setupBalanceRoutes(databaseClient *sql.DB) balance.BalanceRoutes {
	// setup repository
	repository := balanceRepository.NewBalanceRepository(databaseClient)

	// setup Use Cases (services)
	createUseCase := balanceCreate.NewCreateUseCase(repository)
	deleteUseCase := balanceDelete.NewDeleteUseCase(repository)
	findUseCase := balanceFind.NewFindUseCase(repository)
	listUseCase := balanceList.NewListUseCase(repository)
	updateUseCase := balanceUpdate.NewUpdateUseCase(repository)

	// setup router handlers
	balanceHandler := balance.NewBalanceHandler(createUseCase, listUseCase, findUseCase, updateUseCase, deleteUseCase)

	// setup routes
	balanceRoutes := balance.NewBalanceRoutes(balanceHandler)

	return balanceRoutes
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
			logger.Infof("%6s: %s \n", route.Method, route.Path)
		}
	}
}
