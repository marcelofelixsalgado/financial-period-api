package user

import (
	"marcelofelixsalgado/financial-period-api/api/controllers"
	"net/http"
)

var userBasepath = "/v1/users"

type UserRoutes struct {
	userHandler IUserHandler
}

func NewUserRoutes(userHandler IUserHandler) UserRoutes {
	return UserRoutes{
		userHandler: userHandler,
	}
}

func (userRoutes *UserRoutes) UserRouteMapping() []controllers.Route {

	return []controllers.Route{
		{
			URI:                    userBasepath,
			Method:                 http.MethodPost,
			Function:               userRoutes.userHandler.CreateUser,
			RequiresAuthentication: false,
		},
		{
			URI:                    userBasepath,
			Method:                 http.MethodGet,
			Function:               userRoutes.userHandler.ListUsers,
			RequiresAuthentication: true,
		},
		{
			URI:                    userBasepath + "/{id}",
			Method:                 http.MethodGet,
			Function:               userRoutes.userHandler.GetUserById,
			RequiresAuthentication: true,
		},
		{
			URI:                    userBasepath + "/{id}",
			Method:                 http.MethodPut,
			Function:               userRoutes.userHandler.UpdateUser,
			RequiresAuthentication: true,
		},
		{
			URI:                    userBasepath + "/{id}",
			Method:                 http.MethodDelete,
			Function:               userRoutes.userHandler.DeleteUser,
			RequiresAuthentication: true,
		},
	}
}
