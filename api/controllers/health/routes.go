package health

import (
	"marcelofelixsalgado/financial-period-api/api/controllers"
	"net/http"
)

var healthBasepath = "/v1/health"

type HealthRoutes struct {
	healthHandler IHealthHandler
}

func NewHealthRoutes(healthHandler IHealthHandler) HealthRoutes {
	return HealthRoutes{
		healthHandler: healthHandler,
	}
}

func (healthRoutes *HealthRoutes) HealthRouteMapping() (string, []controllers.Route) {
	return healthBasepath, []controllers.Route{
		{
			URI:                    "",
			Method:                 http.MethodGet,
			Function:               healthRoutes.healthHandler.Health,
			RequiresAuthentication: false,
		},
	}
}
