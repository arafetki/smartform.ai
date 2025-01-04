package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/arafetki/smartform.ai/backend/internals/api/handlers"
	"github.com/arafetki/smartform.ai/backend/internals/api/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterHandlers(e *echo.Echo, handler *handlers.Handler) {

	e.HTTPErrorHandler = customHttpErrorHandler

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[ECHO] ${time_rfc3339} | ${method} | ${uri} | ${status} | ${id} | ${latency_human} | ${remote_ip} | ${user_agent} | error: ${error}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 30 * time.Second,
	}))

	e.GET("/health", handler.HealthCheck)

	v1 := e.Group("/v1")
	{
		v1.Use(middlewares.Authenticate)

		v1.GET("/users", handler.ListAllUsers)
		v1.GET("/users/:id", handler.FetchUser, middlewares.RequireAuthenticatedUser)
		v1.GET("/users/:userId/forms", handler.ListFormsForUser, middlewares.RequireAuthenticatedUser)
		v1.POST("/users/webhook", handler.UserWebhook)

		v1.POST("/forms", handler.CreateForm, middlewares.RequireAuthenticatedUser)
		v1.GET("/forms/:id", handler.FetchFormData)
		v1.DELETE("/forms", handler.DeleteForms, middlewares.RequireAuthenticatedUser)
	}
}

func customHttpErrorHandler(err error, c echo.Context) {

	if c.Response().Committed {
		return
	}
	code := http.StatusInternalServerError
	var message any = "The server encountered a problem and could not process your request"
	if httpError, ok := err.(*echo.HTTPError); ok {
		code = httpError.Code
		switch code {
		case http.StatusNotFound:
			message = "The requested resource could not be found."
		case http.StatusMethodNotAllowed:
			message = fmt.Sprintf("The %s method is not supported for this resource", c.Request().Method)
		default:
			message = httpError.Message
		}
	}
	if err := c.JSON(code, echo.Map{"message": message}); err != nil {
		c.Logger().Error(err)
	}
}
