package transactiontype

import (
	"net/http"

	"github.com/marcelofelixsalgado/financial-period-api/api/controllers"
)

var basepath = "/v1/transaction_types"

type TransactionTypeRoutes struct {
	transactionTypeHandler ITransactionTypeHandler
}

func NewTransactionTypeRoutes(transactionTypeHandler ITransactionTypeHandler) TransactionTypeRoutes {
	return TransactionTypeRoutes{
		transactionTypeHandler: transactionTypeHandler,
	}
}

func (transactionTypeRoutes *TransactionTypeRoutes) TransactionTypeRouteMapping() (string, []controllers.Route) {

	return basepath, []controllers.Route{
		{
			URI:                    "",
			Method:                 http.MethodGet,
			Function:               transactionTypeRoutes.transactionTypeHandler.ListTransactionTypes,
			RequiresAuthentication: true,
		},
		{
			URI:                    "/:code",
			Method:                 http.MethodGet,
			Function:               transactionTypeRoutes.transactionTypeHandler.GetTransactionTypeByCode,
			RequiresAuthentication: true,
		},
	}
}
