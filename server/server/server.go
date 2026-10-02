package server

import (
	"log/slog"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogHost:      true,
		LogLatency:   true,
		LogURI:       true,
		LogURIPath:   true,
		LogRequestID: true,
		LogMethod:    true,
		LogStatus:    true,
		LogRemoteIP:  true,
		LogValuesFunc: func(c *echo.Context, v middleware.RequestLoggerValues) error {
			slog.Info("Request Info",
				"Host", v.Host,
				"Uri", v.URI,
				"UriPath", v.URIPath,
				"Method", v.Method,
				"Status", v.Status,
				"RequestIP", v.RemoteIP,
				"Latency", v.Latency,
			)
			return nil
		},
	},
	))
	return e
}
