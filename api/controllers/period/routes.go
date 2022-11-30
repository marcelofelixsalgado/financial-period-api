package period

import (
	"marcelofelixsalgado/financial-period-api/api/controllers"
	"net/http"
)

var periodBasepath = "/v1/periods"

type PeriodRoutes struct {
	periodHandler IPeriodHandler
}

func NewPeriodRoutes(periodHandler IPeriodHandler) PeriodRoutes {
	return PeriodRoutes{
		periodHandler: periodHandler,
	}
}

func (periodRoutes *PeriodRoutes) PeriodRouteMapping() []controllers.Route {

	return []controllers.Route{
		{
			URI:                    periodBasepath,
			Method:                 http.MethodPost,
			Function:               periodRoutes.periodHandler.CreatePeriod,
			RequiresAuthentication: false,
		},
		{
			URI:                    periodBasepath,
			Method:                 http.MethodGet,
			Function:               periodRoutes.periodHandler.ListPeriods,
			RequiresAuthentication: false,
		},
		{
			URI:                    periodBasepath + "/{id}",
			Method:                 http.MethodGet,
			Function:               periodRoutes.periodHandler.GetPeriodById,
			RequiresAuthentication: false,
		},
		{
			URI:                    periodBasepath + "/{id}",
			Method:                 http.MethodPut,
			Function:               periodRoutes.periodHandler.UpdatePeriod,
			RequiresAuthentication: false,
		},
		{
			URI:                    periodBasepath + "/{id}",
			Method:                 http.MethodDelete,
			Function:               periodRoutes.periodHandler.DeletePeriod,
			RequiresAuthentication: false,
		},
	}
}
