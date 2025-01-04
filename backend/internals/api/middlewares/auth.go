package middlewares

import (
	"net/http"

	"github.com/arafetki/smartform.ai/backend/internals/logging"
	"github.com/arafetki/smartform.ai/backend/internals/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Add("Vary", "Authorization")

		userId, err := uuid.Parse("853c4cc4-eabe-49b0-9072-da2176c4af7a")
		if err != nil {
			logging.Logger().Error(err.Error())
			r := utils.ContextSetUser(c.Request(), utils.AnonymousUser)
			c.SetRequest(r)
			return next(c)

		}

		r := utils.ContextSetUser(c.Request(), &utils.DummyUser{
			ID: userId,
		})

		c.SetRequest(r)
		return next(c)
	}
}

func RequireAuthenticatedUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := utils.ContextGetUser(c.Request())
		if user.IsAnonymous() {
			return echo.NewHTTPError(http.StatusUnauthorized, "You must be authenticated to access this resource")
		}

		return next(c)
	}
}
