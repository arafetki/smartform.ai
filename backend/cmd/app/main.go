package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/arafetki/smartform.ai/backend/internals/api"
	"github.com/arafetki/smartform.ai/backend/internals/api/handlers"
	"github.com/arafetki/smartform.ai/backend/internals/api/middlewares"
	"github.com/arafetki/smartform.ai/backend/internals/api/router"
	"github.com/arafetki/smartform.ai/backend/internals/app"
	"github.com/arafetki/smartform.ai/backend/internals/config"
	"github.com/arafetki/smartform.ai/backend/internals/db"
	"github.com/arafetki/smartform.ai/backend/internals/env"
	"github.com/arafetki/smartform.ai/backend/internals/logging"
	"github.com/arafetki/smartform.ai/backend/internals/repository/sqlc"
	"github.com/arafetki/smartform.ai/backend/internals/services"
	"github.com/arafetki/smartform.ai/backend/internals/validator"
)

func init() {
	config.Init()
	logging.Init(logging.Options{
		Debug:  env.GetBool("APP_DEBUG", true),
		Writer: os.Stdout,
	})
}

func main() {
	if err := startApp(); err != nil {
		trace := string(debug.Stack())
		logging.Logger().Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}

func startApp() error {

	cfg := config.Get()

	db, err := db.Init(cfg.Database.URL)
	if err != nil {
		return err
	}
	defer db.Close()

	logging.Logger().Info("Database connection established")

	queries := sqlc.New(db)

	us := services.NewUsersService(queries)
	fs := services.NewFormsService(queries)

	handler := &handlers.Handler{
		UsersService: us,
		FormsService: fs,
	}

	middleware := &middlewares.Middleware{
		UsersService: us,
	}

	validator := validator.New()
	app := app.New(cfg, validator)

	router.RegisterHandlers(app.Router, handler, middleware)

	server := api.NewServer(app.Router, &api.ServerOptions{
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	})

	return server.Start(fmt.Sprintf(":%d", cfg.Server.Port))
}
