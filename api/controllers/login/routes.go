package login

import (
	"net/http"

	"github.com/marcelofelixsalgado/financial-period-api/api/controllers"
)

var userBasepath = "/v1/login"

type LoginRoutes struct {
	userCredentialsHandler ILoginHandler
}

func NewLoginRoutes(userCredentialsHandler ILoginHandler) LoginRoutes {
	return LoginRoutes{
		userCredentialsHandler: userCredentialsHandler,
	}
}

func (userCredentialsRoutes *LoginRoutes) LoginRouteMapping() (string, []controllers.Route) {

	return userBasepath, []controllers.Route{
		{
			URI:                    "",
			Method:                 http.MethodPost,
			Function:               userCredentialsRoutes.userCredentialsHandler.Login,
			RequiresAuthentication: false,
		},
	}
}
