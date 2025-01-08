package middleware

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (m *Middleware) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Add("Vary", "Authorization")
		userId, err := uuid.Parse("b8c746d6-8e40-44d7-ab14-37ef341fd5a7")
		if err != nil {
			m.logger.Error(err.Error())
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid userID")
		}
		c.Set("user", userId)
		return next(c)
	}
}
func (m *Middleware) RequireAuthenticatedUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Get("user")
		if userID == nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "You must be authenticated to access this resource")
		}
		return next(c)
	}
}
