package handlers

import (
	"errors"
	"net/http"

	"github.com/arafetki/smartform.ai/backend/internals/logging"
	"github.com/arafetki/smartform.ai/backend/internals/services"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *Handler) FetchUserData(c echo.Context) error {

	var input struct {
		ID uuid.UUID `param:"id"`
	}
	if err := c.Bind(&input); err != nil {
		logging.Logger().Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	user, err := h.UsersService.GetUserByID(input.ID)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		logging.Logger().Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": user})
}

func (h *Handler) UserWebhook(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
