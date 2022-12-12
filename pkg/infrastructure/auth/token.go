package auth

import (
	"errors"
	"fmt"
	"marcelofelixsalgado/financial-period-api/configs"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(userID string) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["Authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	jwtToken, err := token.SignedString(configs.SecretKey)
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
	return configs.SecretKey, nil
}

func ExtractUserId(r *http.Request) (string, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, getVerificationKey)
	if err != nil {
		return "", err
	}
	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//userID, err := strconv.ParseUint(fmt.Sprintf("%.f", permissions["userId"]), 10, 64)
		return permissions["userId"].(string), nil
	}
	return "", errors.New("ivalid token")
}