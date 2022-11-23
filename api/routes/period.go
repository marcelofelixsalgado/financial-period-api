package routes

import (
	"marcelofelixsalgado/financial-period-api/api/controllers"
	"net/http"
)

var periodBasepath = "/v1/periods"

var periodRoutes = []Route{
	{
		URI:                    periodBasepath,
		Method:                 http.MethodPost,
		Function:               controllers.CreatePeriod,
		RequiresAuthentication: false,
	},
	{
		URI:                    periodBasepath,
		Method:                 http.MethodGet,
		Function:               controllers.ListPeriods,
		RequiresAuthentication: false,
	},
	{
		URI:                    periodBasepath + "/{id}",
		Method:                 http.MethodGet,
		Function:               controllers.GetPeriodById,
		RequiresAuthentication: false,
	},
	{
		URI:                    periodBasepath + "/{id}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdatePeriod,
		RequiresAuthentication: false,
	},
	{
		URI:                    periodBasepath + "/{id}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeletePeriod,
		RequiresAuthentication: false,
	},
}
