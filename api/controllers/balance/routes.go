package balance

import (
	"marcelofelixsalgado/financial-period-api/api/controllers"
	"net/http"
)

var balanceBasePath = "/v1/balances"

type BalanceRoutes struct {
	balanceHandler IBalanceHandler
}

func NewBalanceRoutes(balanceHandler IBalanceHandler) BalanceRoutes {
	return BalanceRoutes{
		balanceHandler: balanceHandler,
	}
}

func (balanceRoutes *BalanceRoutes) BalanceRouteMapping() (string, []controllers.Route) {

	return balanceBasePath, []controllers.Route{
		{
			URI:                    "",
			Method:                 http.MethodPost,
			Function:               balanceRoutes.balanceHandler.CreateBalance,
			RequiresAuthentication: true,
		},
		{
			URI:                    "",
			Method:                 http.MethodGet,
			Function:               balanceRoutes.balanceHandler.ListBalances,
			RequiresAuthentication: true,
		},
		{
			URI:                    "/:id",
			Method:                 http.MethodGet,
			Function:               balanceRoutes.balanceHandler.GetBalanceById,
			RequiresAuthentication: true,
		},
		{
			URI:                    "/:id",
			Method:                 http.MethodPut,
			Function:               balanceRoutes.balanceHandler.UpdateBalance,
			RequiresAuthentication: true,
		},
		{
			URI:                    "/:id",
			Method:                 http.MethodDelete,
			Function:               balanceRoutes.balanceHandler.DeleteBalance,
			RequiresAuthentication: true,
		},
	}
}
