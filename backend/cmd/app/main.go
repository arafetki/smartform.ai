package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/arafetki/smartform.ai/backend/internals/api"
	"github.com/arafetki/smartform.ai/backend/internals/api/handlers"
	"github.com/arafetki/smartform.ai/backend/internals/api/router"
	"github.com/arafetki/smartform.ai/backend/internals/config"
	"github.com/arafetki/smartform.ai/backend/internals/db"
	"github.com/arafetki/smartform.ai/backend/internals/logging"
	"github.com/arafetki/smartform.ai/backend/internals/repository/sqlc"
	"github.com/arafetki/smartform.ai/backend/internals/services"
	"github.com/arafetki/smartform.ai/backend/internals/validator"
	"github.com/labstack/echo/v4"
)

func init() {
	config.Init()
	logging.Init(logging.Options{
		Debug:  config.Get().Application.Debug,
		Writer: os.Stdout,
	})
}

func main() {
	if err := runApp(); err != nil {
		trace := string(debug.Stack())
		logging.Logger().Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}

func runApp() error {
	cfg := config.Get()

	e := echo.New()
	e.Validator = validator.New()

	db, err := db.Init(cfg.Database.URL)
	if err != nil {
		return err
	}
	defer db.Close()
	logging.Logger().Info("Database connection established")

	queries := sqlc.New(db)

	handler := &handlers.Handler{
		UsersService: services.NewUsersService(queries),
		FormsService: services.NewFormsService(queries),
	}
	router.RegisterHandlers(e, handler)

	server := api.NewServer(e, &api.ServerOptions{
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	})

	return server.Start(fmt.Sprintf(":%d", cfg.Server.Port))
}
