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
		SettingsID  int16           `json:"settings_id" validate:"required"`
		Title       string          `json:"title" validate:"required"`
		Description pgtype.Text     `json:"description"`
		Fields      json.RawMessage `json:"fields" validate:"required"`
		Published   bool            `json:"published"`
	}

	if err := c.Bind(&input); err != nil {
		logging.Logger().Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if err := c.Validate(input); err != nil {
		logging.Logger().Error(err.Error())
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
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
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusCreated)
}

func (h *Handler) FetchFormData(c echo.Context) error {
	var input struct {
		ID uuid.UUID `param:"id"`
	}
	if err := c.Bind(&input); err != nil {
		logging.Logger().Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	data, err := h.FormsService.GetFormWithSettings(input.ID)
	if err != nil {
		if errors.Is(err, services.ErrFormNotFound) {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		logging.Logger().Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": data})
}

func (h *Handler) FetchFormsForUser(c echo.Context) error {

	authenticatedUser := utils.ContextGetUser(c.Request())

	forms, err := h.FormsService.ListFormsForUser(authenticatedUser.ID)
	if err != nil {
		logging.Logger().Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": forms})
}

func (h *Handler) DeleteFormsInBatch(c echo.Context) error {
	var input struct {
		Ids []uuid.UUID `json:"ids"`
	}
	if err := c.Bind(&input); err != nil {
		logging.Logger().Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	authenticatedUser := utils.ContextGetUser(c.Request())

	err := h.FormsService.DeleteForms(sqlc.DeleteFormsByOwnerParams{Ids: input.Ids, UserID: authenticatedUser.ID})
	if err != nil {
		switch {
		case errors.Is(err, services.ErrNoFormsDeleted):
			return echo.NewHTTPError(http.StatusNotFound)
		default:
			logging.Logger().Error(err.Error())
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}
	return c.NoContent(http.StatusNoContent)
}
