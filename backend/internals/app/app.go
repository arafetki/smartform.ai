package app

import (
	"github.com/arafetki/smartform.ai/backend/internals/config"
	"github.com/arafetki/smartform.ai/backend/internals/validator"
	"github.com/labstack/echo/v4"
)

type Application struct {
	Router *echo.Echo
}

func New(cfg config.Config, validator *validator.Validator) *Application {
	e := echo.New()
	e.Debug = cfg.Application.Debug
	e.Validator = validator
	return &Application{
		Router: e,
	}
}
