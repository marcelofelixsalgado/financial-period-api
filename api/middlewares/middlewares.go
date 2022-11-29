package middlewares

import (
	"log"
	"marcelofelixsalgado/financial-period-api/api/auth"
	"marcelofelixsalgado/financial-period-api/api/responses"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidateToken(r); err != nil {
			log.Printf("Token validation error: %v", err)
			responses.JSONErrorByCode(w, responses.NotAuthorized)
			return
		}
		next(w, r)
	}
}

func ResponseFormatMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
