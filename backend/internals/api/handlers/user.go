package handlers

import (
	"errors"
	"net/http"

	"github.com/arafetki/smartform.ai/backend/internals/logging"
	"github.com/arafetki/smartform.ai/backend/internals/services"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *Handler) FetchUser(c echo.Context) error {

	var input struct {
		ID uuid.UUID `param:"id"`
	}
	if err := c.Bind(&input); err != nil {
		logging.Logger().Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "The request could not be understood by the server due to malformed syntax or incorrect parameter type")
	}

	user, err := h.UsersService.GetUserByID(input.ID)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "The requested resource could not be found.")
		}
		logging.Logger().Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "the server encountered a problem and could not process your request.")
	}

	return c.JSON(http.StatusOK, echo.Map{"data": user})
}

func (h *Handler) ListAllUsers(c echo.Context) error {

	users, err := h.UsersService.ListAllUsers()
	if err != nil {
		logging.Logger().Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "the server encountered a problem and could not process your request.")
	}
	return c.JSON(http.StatusOK, echo.Map{"data": users})
}

func (h *Handler) UserWebhook(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) ListFormsForUser(c echo.Context) error {
	var input struct {
		UserId uuid.UUID `param:"userId"`
	}
	if err := c.Bind(&input); err != nil {
		logging.Logger().Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "The request could not be understood by the server due to malformed syntax or incorrect parameter type")
	}

	forms, err := h.FormsService.ListFormsForUser(input.UserId)
	if err != nil {
		logging.Logger().Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "the server encountered a problem and could not process your request.")
	}

	return c.JSON(http.StatusOK, echo.Map{"data": forms})
}
