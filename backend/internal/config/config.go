package config

import (
	"fmt"
	"time"

	"github.com/arafetki/smartform.ai/backend/internal/env"
)

type Config struct {
	App struct {
		Name        string
		Description string
		Version     string
		Env         string
		Debug       bool
	}
	Server struct {
		Addr           string
		ReadTimeout    time.Duration
		WriteTimeout   time.Duration
		ShutdownPeriod time.Duration
	}
	Database struct {
		Dsn         string
		Automigrate bool
	}
	JWT struct {
		Secret string
	}
}

func Init() Config {
	var cfg Config

	cfg.App.Name = "SmartForm AI"
	cfg.App.Description = "JSON RESTful API"
	cfg.App.Version = "0.1.0"
	cfg.App.Env = env.GetString("APP_ENV", "development")
	cfg.App.Debug = env.GetBool("APP_DEBUG", true)

	cfg.Server.Addr = fmt.Sprintf(":%d", env.GetInt("SERVER_HTTP_PORT", 8080))
	cfg.Server.ReadTimeout = env.GetDuration("SERVER_READ_TIMEOUT", 10*time.Second)
	cfg.Server.WriteTimeout = env.GetDuration("SERVER_WRITE_TIMEOUT", 20*time.Second)
	cfg.Server.ShutdownPeriod = env.GetDuration("SERVER_SHUTDOWN_PERIOD", 30*time.Second)

	cfg.Database.Dsn = env.GetString("DATABASE_DSN", "postgres:postgres@localhost:5432/smartform?sslmode=disable")
	cfg.JWT.Secret = env.GetString("JWT_SECRET_KEY", "secret")

	if cfg.App.Env == "development" {
		cfg.Database.Automigrate = true
	} else {
		cfg.Database.Automigrate = false
	}

	return cfg
}
