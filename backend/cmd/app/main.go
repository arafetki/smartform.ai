package main

import (
	"fmt"
	"os"

	"github.com/arafetki/smartform.ai/backend/internals/api"
	"github.com/arafetki/smartform.ai/backend/internals/api/handlers"
	"github.com/arafetki/smartform.ai/backend/internals/api/router"
	"github.com/arafetki/smartform.ai/backend/internals/config"
	"github.com/arafetki/smartform.ai/backend/internals/db"
	"github.com/arafetki/smartform.ai/backend/internals/logging"
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
	cfg := config.Get()

	e := echo.New()
	e.Validator = validator.New()

	db, err := db.Init(cfg.Database.URL)
	if err != nil {
		logging.Logger().Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()
	logging.Logger().Info("Database connection established")

	handler := &handlers.Handler{}
	router.RegisterHandlers(e, handler)

	server := api.NewServer(e)
	if err := server.Start(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
		logging.Logger().Error(err.Error())
		os.Exit(1)
	}
}
