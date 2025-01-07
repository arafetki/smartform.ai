package app

import (
	"fmt"
	"net/http"
	"sync"

	rest "github.com/arafetki/smartform.ai/backend/internal/app/api"
	"github.com/arafetki/smartform.ai/backend/internal/app/api/handler"
	"github.com/arafetki/smartform.ai/backend/internal/config"
	"github.com/arafetki/smartform.ai/backend/internal/logging"
	"github.com/arafetki/smartform.ai/backend/internal/service"
	"github.com/labstack/echo/v4"
)

type application struct {
	echo   *echo.Echo
	logger *logging.Logger
	cfg    config.Config
	svc    *service.Service
	wg     sync.WaitGroup
}

func New(cfg config.Config, logger *logging.Logger, svc *service.Service) *application {
	app := &application{
		echo:   echo.New(),
		logger: logger,
		cfg:    cfg,
		svc:    svc,
	}

	app.configure()

	return app
}

func (app *application) Run() error {

	// Register rest endpoints
	rest.Routes(app.echo, handler.New(app.logger, app.cfg, app.svc))

	// Start http server
	return app.start()
}

func (app *application) configure() {
	app.echo.HideBanner = true
	app.echo.HidePort = true
	app.echo.Debug = app.cfg.App.Debug
	app.echo.Server.ReadTimeout = app.cfg.Server.ReadTimeout
	app.echo.Server.WriteTimeout = app.cfg.Server.WriteTimeout
	app.echo.HTTPErrorHandler = handleErrors(app.logger)
}

func handleErrors(logger *logging.Logger) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}
		code := http.StatusInternalServerError
		var message any = "The server encountered a problem and could not process your request."
		if httpError, ok := err.(*echo.HTTPError); ok {
			code = httpError.Code
			switch code {
			case http.StatusNotFound:
				message = "The requested resource could not be found."
			case http.StatusMethodNotAllowed:
				message = fmt.Sprintf("The %s method is not supported for this resource.", c.Request().Method)
			case http.StatusBadRequest:
				message = "The request could not be understood or was missing required parameters."
			case http.StatusInternalServerError:
				message = "The server encountered a problem and could not process your request."
			case http.StatusUnprocessableEntity:
				message = "The request could not be processed due to invalid input."
			default:
				message = httpError.Message
			}
		} else {
			logger.Error(err.Error())
		}
		if err := c.JSON(code, echo.Map{"message": message}); err != nil {
			logger.Error(err.Error())
		}
	}
}
