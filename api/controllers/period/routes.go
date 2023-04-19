package period

import (
	"net/http"

	"github.com/marcelofelixsalgado/financial-period-api/api/controllers"
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

func (periodRoutes *PeriodRoutes) PeriodRouteMapping() (string, []controllers.Route) {

	return periodBasepath, []controllers.Route{
		{
			URI:                    "",
			Method:                 http.MethodPost,
			Function:               periodRoutes.periodHandler.CreatePeriod,
			RequiresAuthentication: true,
		},
		{
			URI:                    "",
			Method:                 http.MethodGet,
			Function:               periodRoutes.periodHandler.ListPeriods,
			RequiresAuthentication: true,
		},
		{
			URI:                    "/:id",
			Method:                 http.MethodGet,
			Function:               periodRoutes.periodHandler.GetPeriodById,
			RequiresAuthentication: true,
		},
		{
			URI:                    "/:id",
			Method:                 http.MethodPut,
			Function:               periodRoutes.periodHandler.UpdatePeriod,
			RequiresAuthentication: true,
		},
		{
			URI:                    "/:id",
			Method:                 http.MethodDelete,
			Function:               periodRoutes.periodHandler.DeletePeriod,
			RequiresAuthentication: true,
		},
	}
}
