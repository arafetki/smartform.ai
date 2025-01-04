package router

import (
	"github.com/arafetki/smartform.ai/backend/internals/api/handlers"
	"github.com/arafetki/smartform.ai/backend/internals/api/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterHandlers(r *echo.Echo, h *handlers.Handler, m *middlewares.Middleware) {

	r.HTTPErrorHandler = h.CustomHttpErrorHandler

	r.Use(middleware.CORS())
	r.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[ECHO] ${time_rfc3339} | ${method} | ${uri} | ${status} | ${id} | ${latency_human} | ${remote_ip} | ${user_agent} | error: ${error}\n",
	}))
	r.Use(middleware.Recover())
	r.Use(middleware.RequestID())

	r.GET("/health", h.HealthCheck)

	v1 := r.Group("/v1")
	{
		v1.Use(m.Authenticate)

		v1.GET("/users/:id", h.FetchUserData, m.RequireAuthenticatedUser)
		v1.GET("/users/:id/forms", h.FetchFormsForUser, m.RequireAuthenticatedUser)
		v1.POST("/users/webhook", h.UserWebhook)

		v1.POST("/forms", h.CreateForm, m.RequireAuthenticatedUser)
		v1.GET("/forms", h.FetchFormsForUser, m.RequireAuthenticatedUser)
		v1.DELETE("/forms", h.DeleteFormsInBatch, m.RequireAuthenticatedUser)
		v1.GET("/forms/:id", h.FetchFormData)
	}
}
