package logger

import (
	"context"
	"marcelofelixsalgado/financial-period-api/settings"
	"marcelofelixsalgado/financial-period-api/version"
	"os"

	custom_context "marcelofelixsalgado/financial-period-api/api/context"

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

// GetLoggerTemplate - Returns a log template
func (l LogParameters) GetLoggerTemplate(ctx context.Context) *logrus.Entry {
	var operation = map[string]interface{}{
		"module": l.Module,
		"name":   l.Action,
	}
	entry := GetLogger().WithFields(logrus.Fields{
		"service.name": settings.Config.AppName,
		"operation":    operation,
		"trace.id":     ctx.Value("messageIDKey"),
	})

	if l.EchoContext != nil {
		entry.Data["http"] = custom_context.ContextRequestHTTP(l.EchoContext)
	}

	return entry
}

// GetLogger - Returns an instance of logger
func GetLogger() *logrus.Entry {
	if entry == nil {
		initLogger()
	}

	return entry.WithFields(logrus.Fields{
		"environment": settings.Config.Environment,
		"app.host":    hostname,
		"app.version": version.Version,
		"type":        "json",
	})
}

func initLogger() {
	entry = logrus.New()
	hostname, _ = os.Hostname()
}
