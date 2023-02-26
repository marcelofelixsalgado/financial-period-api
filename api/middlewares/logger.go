package middlewares

import (
	"log"
	"marcelofelixsalgado/financial-period-api/settings"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func Logger() echo.MiddlewareFunc {

	file, err := os.OpenFile(settings.Config.LogAccessFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log := logrus.New()

	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.SetFormatter(&logrus.JSONFormatter{})
			log.SetOutput(file)

			if values.Error == nil {
				log.WithFields(logrus.Fields{
					"method": c.Request().Method,
					"URI":    values.URI,
					"path":   c.Path(),
					"status": values.Status,
				}).Info("request")
			} else {
				log.WithFields(logrus.Fields{
					"method": c.Request().Method,
					"URI":    values.URI,
					"path":   c.Path(),
					"status": values.Status,
					"error":  values.Error,
				}).Error("request error")
			}
			return nil
		},
	})
}
