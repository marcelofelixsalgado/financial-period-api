package auth

import (
	"errors"
	"fmt"
	"marcelofelixsalgado/financial-period-api/settings"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userId string, tenantId string) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["Authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userId
	permissions["tenantId"] = tenantId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	jwtToken, err := token.SignedString(settings.Config.SecretKey)
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}

func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, getVerificationKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("invalid token")
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	// Bearer 123
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func getVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signature method! %v", token.Header["alg"])
	}
	return settings.Config.SecretKey, nil
}

func extractClaims(claim string, r *http.Request) (string, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, getVerificationKey)
	if err != nil {
		return "", err
	}
	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return permissions[claim].(string), nil
	}
	return "", errors.New("ivalid token")
}

func ExtractUserId(r *http.Request) (string, error) {
	return extractClaims("userId", r)
}

func ExtractTenantId(r *http.Request) (string, error) {
	return extractClaims("tenantId", r)
}
