package controllers

import "github.com/labstack/echo/v4"

type Route struct {
	URI                    string
	Method                 string
	Function               func(c echo.Context) error
	RequiresAuthentication bool
}
