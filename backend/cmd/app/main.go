package main

import (
	"fmt"

	"github.com/arafetki/smartform.ai/backend/internals/api"
	"github.com/arafetki/smartform.ai/backend/internals/api/handlers"
	"github.com/arafetki/smartform.ai/backend/internals/api/router"
	"github.com/arafetki/smartform.ai/backend/internals/config"
	"github.com/arafetki/smartform.ai/backend/internals/db"
	"github.com/arafetki/smartform.ai/backend/internals/validator"
	"github.com/labstack/echo/v4"
)

func init() {
	config.Init()
}

func main() {
	e := echo.New()
	cfg := config.Get()
	e.Debug = cfg.Debug
	e.Validator = validator.New()
	db, err := db.Init(cfg.Database.URL)
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()
	e.Logger.Info("Database connection established")
	handler := &handlers.Handler{}
	router.RegisterHandlers(e, handler)
	server := api.NewServer(e)
	if err := server.Start(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		e.Logger.Fatal(err)
	}
}
