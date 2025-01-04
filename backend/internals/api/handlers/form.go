package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/arafetki/smartform.ai/backend/internals/logging"
	"github.com/arafetki/smartform.ai/backend/internals/repository/sqlc"
	"github.com/arafetki/smartform.ai/backend/internals/services"
	"github.com/arafetki/smartform.ai/backend/internals/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateForm(c echo.Context) error {
	var input struct {
		SettingsID  pgtype.Int2     `json:"settings_id"`
		Title       string          `json:"title"`
		Description pgtype.Text     `json:"description"`
		Fields      json.RawMessage `json:"fields"`
		Published   bool            `json:"published"`
	}

	if err := c.Bind(&input); err != nil {
		logging.Logger().Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "The request could not be understood by the server due to malformed syntax or incorrect parameter type")
	}

	authenticatedUser := utils.ContextGetUser(c.Request())

	params := sqlc.CreateFormParams{
		UserID:      authenticatedUser.ID,
		SettingsID:  input.SettingsID,
		Title:       input.Title,
		Description: input.Description,
		Fields:      input.Fields,
		Published:   input.Published,
	}

	err := h.FormsService.CreateForm(params)
	if err != nil {
		logging.Logger().Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "the server encountered a problem and could not process your request.")
	}
	return c.NoContent(http.StatusCreated)
}

func (h *Handler) FetchFormData(c echo.Context) error {
	var input struct {
		ID uuid.UUID `param:"id"`
	}
	if err := c.Bind(&input); err != nil {
		logging.Logger().Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "The request could not be understood by the server due to malformed syntax or incorrect parameter type")
	}
	data, err := h.FormsService.GetFormWithSettings(input.ID)
	if err != nil {
		if errors.Is(err, services.ErrFormNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "The requested resource could not be found.")
		}
		logging.Logger().Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "the server encountered a problem and could not process your request.")
	}

	return c.JSON(http.StatusOK, echo.Map{"data": data})
}

func (h *Handler) DeleteForms(c echo.Context) error {
	var input struct {
		IDS []uuid.UUID `json:"ids"`
	}
	if err := c.Bind(&input); err != nil {
		logging.Logger().Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "The request could not be understood by the server due to malformed syntax or incorrect parameter type")
	}

	authenticatedUser := utils.ContextGetUser(c.Request())

	err := h.FormsService.DeleteForms(sqlc.DeleteFormsByOwnerParams{ID: input.IDS, UserID: authenticatedUser.ID})
	if err != nil {
		switch {
		case errors.Is(err, services.ErrNoFormsDeleted):
			return echo.NewHTTPError(http.StatusNotFound, "The requested resource(s) could not be found.")
		default:
			logging.Logger().Error(err.Error())
			return echo.NewHTTPError(http.StatusInternalServerError, "the server encountered a problem and could not process your request.")
		}
	}
	return c.NoContent(http.StatusNoContent)
}
