package logger

import (
	"github.com/go-chi/httplog/v2"
	"lobinv-server/config"
)

func New(config config.Config) *httplog.Logger {
	logger := httplog.NewLogger("lobinv-server", httplog.Options{
		LogLevel: config.Server.LogLevel,
		Concise:  true,
	})

	return logger
}
