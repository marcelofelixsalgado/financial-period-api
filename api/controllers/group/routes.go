package group

import (
	"marcelofelixsalgado/financial-period-api/api/controllers"
	"net/http"
)

var groupBasepath = "/v1/groups"

type GroupRoutes struct {
	groupHandler IGroupHandler
}

func NewGroupRoutes(groupHandler IGroupHandler) GroupRoutes {
	return GroupRoutes{
		groupHandler: groupHandler,
	}
}

func (groupRoutes *GroupRoutes) GroupRouteMapping() (string, []controllers.Route) {

	return groupBasepath, []controllers.Route{
		{
			URI:                    "",
			Method:                 http.MethodPost,
			Function:               groupRoutes.groupHandler.CreateGroup,
			RequiresAuthentication: true,
		},
		{
			URI:                    "",
			Method:                 http.MethodGet,
			Function:               groupRoutes.groupHandler.ListGroups,
			RequiresAuthentication: true,
		},
		{
			URI:                    "/:id",
			Method:                 http.MethodGet,
			Function:               groupRoutes.groupHandler.GetGroupById,
			RequiresAuthentication: true,
		},
		{
			URI:                    "/:id",
			Method:                 http.MethodPut,
			Function:               groupRoutes.groupHandler.UpdateGroup,
			RequiresAuthentication: true,
		},
		{
			URI:                    "/:id",
			Method:                 http.MethodDelete,
			Function:               groupRoutes.groupHandler.DeleteGroup,
			RequiresAuthentication: true,
		},
	}
}
