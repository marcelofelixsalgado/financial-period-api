package routes

import (
	"marcelofelixsalgado/financial-period-api/api/controllers"
	"net/http"
)

var healthBasepath = "/v1/health"

var healthRoutes = []Route{
	{
		URI:                    healthBasepath,
		Method:                 http.MethodGet,
		Function:               controllers.Health,
		RequiresAuthentication: false,
	},
}
