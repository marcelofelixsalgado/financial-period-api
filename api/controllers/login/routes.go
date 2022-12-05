package login

import (
	"marcelofelixsalgado/financial-period-api/api/controllers"
	"net/http"
)

var loginBasepath = "/v1/login"

type LoginRoutes struct {
	loginHandler ILoginHandler
}

func NewLoginRoutes(loginHandler ILoginHandler) LoginRoutes {
	return LoginRoutes{
		loginHandler: loginHandler,
	}
}

func (loginRoutes *LoginRoutes) LoginRouteMapping() []controllers.Route {

	return []controllers.Route{
		{
			URI:                    loginBasepath,
			Method:                 http.MethodPost,
			Function:               loginRoutes.loginHandler.Login,
			RequiresAuthentication: false,
		},
	}
}
