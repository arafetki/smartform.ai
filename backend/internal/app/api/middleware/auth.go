package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/arafetki/smartform.ai/backend/internal/db/sqlc"
	"github.com/arafetki/smartform.ai/backend/internal/jwt"
	"github.com/arafetki/smartform.ai/backend/internal/service"
	"github.com/labstack/echo/v4"
)

var anonymousUser = &sqlc.User{}

func (m *Middleware) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Add("Vary", "Authorization")
		authorizationHeader := c.Request().Header.Get("Authorization")

		if authorizationHeader == "" {
			c.Set("user", anonymousUser)
			return next(c)
		}

		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or missing authentication token.")
		}

		tokenString := headerParts[1]
		userId, err := jwt.HMACCheck(tokenString, m.cfg.JWT.Secret)
		if err != nil {
			m.logger.Error(err.Error())
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or missing authentication token.")
		}

		user, err := m.service.Users.GetOne(c.Request().Context(), userId)
		if err != nil {
			if errors.Is(err, service.ErrUserNotFound) {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or missing authentication token.")
			}
			m.logger.Error(err.Error())
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		c.Set("user", user)
		return next(c)
	}
}
func (m *Middleware) RequireAuthenticatedUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, ok := c.Get("user").(*sqlc.User)
		if !ok || user == anonymousUser {
			return echo.NewHTTPError(http.StatusUnauthorized, "You must be authenticated to access this resource.")
		}
		return next(c)
	}
}

func (m *Middleware) RequireVerifiedUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*sqlc.User)
		if !user.IsEmailVerified {
			return echo.NewHTTPError(http.StatusUnauthorized, "You must verify your email to access this resource.")
		}
		return next(c)
	}
}
