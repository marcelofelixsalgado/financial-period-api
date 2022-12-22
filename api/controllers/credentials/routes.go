package credentials

import (
	"marcelofelixsalgado/financial-period-api/api/controllers"
	"net/http"
)

var userBasepath = "/v1"

type UserCredentialsRoutes struct {
	userCredentialsHandler IUserCredentialsHandler
}

func NewUserCredentialsRoutes(userCredentialsHandler IUserCredentialsHandler) UserCredentialsRoutes {
	return UserCredentialsRoutes{
		userCredentialsHandler: userCredentialsHandler,
	}
}

func (userCredentialsRoutes *UserCredentialsRoutes) UserCredentialsRouteMapping() []controllers.Route {

	return []controllers.Route{
		{
			URI:                    userBasepath + "/users/{id}/credentials",
			Method:                 http.MethodPost,
			Function:               userCredentialsRoutes.userCredentialsHandler.CreateUserCredentials,
			RequiresAuthentication: false,
		},
		{
			URI:                    userBasepath + "/users/{id}/credentials",
			Method:                 http.MethodPut,
			Function:               userCredentialsRoutes.userCredentialsHandler.UpdateUserCredentials,
			RequiresAuthentication: true,
		},
		{
			URI:                    userBasepath + "/login",
			Method:                 http.MethodPost,
			Function:               userCredentialsRoutes.userCredentialsHandler.Login,
			RequiresAuthentication: false,
		},
	}
}
