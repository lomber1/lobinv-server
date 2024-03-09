package config

import "log/slog"

type ServerConfig struct {
	Port     int
	LogLevel slog.Level
}

type DatabaseConfig struct {
	Uri string
}

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

func GetConfig() Config {
	return Config{
		Server: ServerConfig{
			Port:     8080,
			LogLevel: slog.LevelError,
		},
		Database: DatabaseConfig{
			Uri: "./database.sqlite",
		},
	}
}
