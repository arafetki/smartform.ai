package api

import (
	"net/http"

	"github.com/arafetki/smartform.ai/backend/internal/app/api/handler"
	"github.com/arafetki/smartform.ai/backend/internal/app/api/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func Routes(r *echo.Echo, h *handler.Handler, m *middleware.Middleware) {

	// Middleware
	r.Use(echoMiddleware.RequestID())
	r.Use(echoMiddleware.LoggerWithConfig(echoMiddleware.LoggerConfig{
		Format: "[ECHO] ${time_rfc3339} | ${method} | ${uri} | ${status} | ${latency_human} | ${remote_ip} | ${user_agent} | error: ${error}\n",
	}))
	r.Use(echoMiddleware.Recover())
	r.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodHead, http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete},
	}))

	// Health checks endpoint
	r.GET("/health", h.HealthCheckHandler)

	// Sub-router
	v1 := r.Group("/v1", m.Authenticate)

	v1.POST("/forms", h.CreateFormHandler, m.RequireAuthenticatedUser, m.RequireVerifiedUser)
	v1.DELETE("/forms/:id", h.DeleteFormHandler, m.RequireAuthenticatedUser, m.RequireVerifiedUser)
}
