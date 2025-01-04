package handlers

import (
	"errors"
	"net/http"

	"github.com/arafetki/smartform.ai/backend/internals/services"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *Handler) FetchUser(c echo.Context) error {

	var input struct {
		ID uuid.UUID `param:"id"`
	}
	if err := c.Bind(&input); err != nil {
		return h.badRequestErrorResponse(c, err)
	}

	user, err := h.UsersService.GetUserByID(input.ID)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			return h.notFoundErrorResponse()
		}
		return h.internalServerErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": user})
}

func (h *Handler) ListAllUsers(c echo.Context) error {

	users, err := h.UsersService.ListAllUsers()
	if err != nil {
		return h.internalServerErrorResponse(c, err)
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
		return h.badRequestErrorResponse(c, err)
	}

	forms, err := h.FormsService.ListFormsForUser(input.UserId)
	if err != nil {
		return h.internalServerErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": forms})
}
