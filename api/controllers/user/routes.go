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

func (userRoutes *UserRoutes) UserRouteMapping() (string, []controllers.Route) {

	return userBasepath, []controllers.Route{
		{
			URI:                    "",
			Method:                 http.MethodPost,
			Function:               userRoutes.userHandler.CreateUser,
			RequiresAuthentication: false,
		},
		{
			URI:                    "",
			Method:                 http.MethodGet,
			Function:               userRoutes.userHandler.ListUsers,
			RequiresAuthentication: true,
		},
		{
			URI:                    "/:id",
			Method:                 http.MethodGet,
			Function:               userRoutes.userHandler.GetUserById,
			RequiresAuthentication: true,
		},
		{
			URI:                    "/:id",
			Method:                 http.MethodPut,
			Function:               userRoutes.userHandler.UpdateUser,
			RequiresAuthentication: true,
		},
		{
			URI:                    "/:id",
			Method:                 http.MethodDelete,
			Function:               userRoutes.userHandler.DeleteUser,
			RequiresAuthentication: true,
		},
		{
			URI:                    "/:id/credentials",
			Method:                 http.MethodPost,
			Function:               userRoutes.userHandler.CreateUserCredentials,
			RequiresAuthentication: false,
		},
		{
			URI:                    "/:id/credentials",
			Method:                 http.MethodPut,
			Function:               userRoutes.userHandler.UpdateUserCredentials,
			RequiresAuthentication: true,
		},
	}
}
