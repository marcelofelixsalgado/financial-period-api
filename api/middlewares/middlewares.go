package middlewares

import (
	"log"
	"marcelofelixsalgado/financial-period-api/api/responses"
	"marcelofelixsalgado/financial-period-api/api/responses/faults"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/auth"

	"github.com/labstack/echo/v4"
)

// func Logger(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
// 		next(w, r)
// 	}
// }

func Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := auth.ValidateToken(c.Request()); err != nil {
			log.Printf("Token validation error: %v", err)
			responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.NotAuthorized)
			return c.JSON(responseMessage.HttpStatusCode, responseMessage)
		}
		return next(c)
	}
}

// func ResponseFormatMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Add("Content-Type", "application/json")
// 		next.ServeHTTP(w, r)
// 	})
// }
