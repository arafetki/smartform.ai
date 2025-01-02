package config

import (
	"time"

	"github.com/arafetki/smartform.ai/backend/internals/env"
)

type Application struct {
	Env   string
	Debug bool
}
type Server struct {
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Database struct {
	URL string
}

type Config struct {
	Application Application
	Server      Server
	Database    Database
}

var cfg *Config

func Init() {

	cfg = &Config{
		Application: Application{
			Debug: env.GetBool("APP_DEBUG", true),
			Env:   env.GetString("APP_ENV", "development"),
		},
		Server: Server{
			Port:         env.GetInt("SERVER_PORT", 8080),
			ReadTimeout:  env.GetDuration("SERVER_READ_TIMEOUT", 15*time.Second),
			WriteTimeout: env.GetDuration("SERVER_WRITE_TIMEOUT", 45*time.Second),
		},
		Database: Database{
			URL: env.GetString("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"),
		},
	}
}

func Get() Config {
	if cfg == nil {
		panic("Configuration is not initialized.")
	}
	return *cfg
}
