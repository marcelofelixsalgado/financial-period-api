package routes

import (
	"marcelofelixsalgado/financial-period-api/api/controllers"
	"marcelofelixsalgado/financial-period-api/api/controllers/balance"
	"marcelofelixsalgado/financial-period-api/api/controllers/category"
	"marcelofelixsalgado/financial-period-api/api/controllers/health"
	"marcelofelixsalgado/financial-period-api/api/controllers/login"
	"marcelofelixsalgado/financial-period-api/api/controllers/period"
	"marcelofelixsalgado/financial-period-api/api/controllers/subcategory"
	"marcelofelixsalgado/financial-period-api/api/controllers/transactiontype"
	"marcelofelixsalgado/financial-period-api/api/controllers/user"
	"marcelofelixsalgado/financial-period-api/api/middlewares"

	"github.com/labstack/echo/v4"
)

type Routes struct {
	loginRoutes           login.LoginRoutes
	userRoutes            user.UserRoutes
	transactionTypeRoutes transactiontype.TransactionTypeRoutes
	categoryRoutes        category.CategoryRoutes
	subCategoryRoutes     subcategory.SubCategoryRoutes
	periodRoutes          period.PeriodRoutes
	balanceRoutes         balance.BalanceRoutes
	healthRoutes          health.HealthRoutes
}

func NewRoutes(loginRoutes login.LoginRoutes, userRoutes user.UserRoutes,
	transactionTypeRoutes transactiontype.TransactionTypeRoutes,
	categoryRoutes category.CategoryRoutes,
	subCategoryRoutes subcategory.SubCategoryRoutes,
	periodRoutes period.PeriodRoutes, balanceRoutes balance.BalanceRoutes,
	healthRoutes health.HealthRoutes) *Routes {
	return &Routes{
		loginRoutes:           loginRoutes,
		userRoutes:            userRoutes,
		transactionTypeRoutes: transactionTypeRoutes,
		categoryRoutes:        categoryRoutes,
		subCategoryRoutes:     subCategoryRoutes,
		periodRoutes:          periodRoutes,
		balanceRoutes:         balanceRoutes,
		healthRoutes:          healthRoutes,
	}
}

func (routes *Routes) RouteMapping(http *echo.Echo) {

	// user credentials routes
	basePath, loginRoutes := routes.loginRoutes.LoginRouteMapping()
	setupRoute(http, basePath, loginRoutes)

	// user routes
	basePath, userRoutes := routes.userRoutes.UserRouteMapping()
	setupRoute(http, basePath, userRoutes)

	// transaction type routes
	basePath, transactionTypeRoutes := routes.transactionTypeRoutes.TransactionTypeRouteMapping()
	setupRoute(http, basePath, transactionTypeRoutes)

	// category routes
	basePath, categoryRoutes := routes.categoryRoutes.CategoryRouteMapping()
	setupRoute(http, basePath, categoryRoutes)

	// subcategory routes
	basePath, subCategoryRoutes := routes.subCategoryRoutes.SubCategoryRouteMapping()
	setupRoute(http, basePath, subCategoryRoutes)

	// period routes
	basePath, periodRoutes := routes.periodRoutes.PeriodRouteMapping()
	setupRoute(http, basePath, periodRoutes)

	// balance routes
	basePath, balanceRoutes := routes.balanceRoutes.BalanceRouteMapping()
	setupRoute(http, basePath, balanceRoutes)

	// health routes
	basePath, healthRoutes := routes.healthRoutes.HealthRouteMapping()
	setupRoute(http, basePath, healthRoutes)

	// return server
}

func setupRoute(http *echo.Echo, basePath string, routes []controllers.Route) {
	group := http.Group(basePath)

	for _, route := range routes {
		if route.RequiresAuthentication {
			group.Add(route.Method, route.URI, route.Function, middlewares.Authenticate)
		} else {
			group.Add(route.Method, route.URI, route.Function)
		}
	}
}
