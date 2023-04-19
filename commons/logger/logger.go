package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/marcelofelixsalgado/financial-period-api/settings"
	"github.com/marcelofelixsalgado/financial-period-api/version"

	"github.com/labstack/echo/v4"

	"github.com/sirupsen/logrus"
)

// LogParameters template
type LogParameters struct {
	Action      string
	Module      string
	EchoContext echo.Context
}

var (
	entry    *logrus.Logger
	hostname string
)

// GetLogger - Returns an instance of logger
func GetLogger() *logrus.Entry {
	if entry == nil {
		initLogger()
	}

	return entry.WithFields(logrus.Fields{
		"environment": settings.Config.Environment,
		"app.host":    hostname,
		"app.version": version.Version,
	})
}

func initLogger() {
	entry = logrus.New()
	hostname, _ = os.Hostname()

	logLevel, err := logrus.ParseLevel(settings.Config.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	entry.SetFormatter(&logrus.JSONFormatter{})
	entry.SetOutput(getLogFile())
	entry.SetLevel(logLevel)
}

func getLogFile() *os.File {

	file, err := os.OpenFile(settings.Config.LogAppFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		os.Exit(1)
	}

	return file
}
