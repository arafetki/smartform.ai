package config

import (
	"log/slog"
	"os"

	"github.com/arafetki/smartform.ai/backend/internals/env"
)

type Config struct {
	Port     int
	Debug    bool
	Database struct {
		URL string
	}
}

var cfg *Config

func Init() {

	cfg = &Config{
		Port:  env.GetInt("PORT", 8080),
		Debug: env.GetBool("APP_DEBUG", true),
		Database: struct{ URL string }{
			URL: env.GetString("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"),
		},
	}
}

func Get() Config {
	if cfg == nil {
		slog.Error("Configuration is not initialized.")
		os.Exit(1)
	}
	return *cfg
}
