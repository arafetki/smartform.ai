package handler

import (
	"encoding/json"
	"net/http"

	"github.com/arafetki/smartform.ai/backend/internal/db/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateFormHandler(c echo.Context) error {

	var input struct {
		Title       string          `json:"title" validate:"required"`
		Description pgtype.Text     `json:"description"`
		Fields      json.RawMessage `json:"fields"`
		Published   bool            `json:"published"`
	}

	if err := c.Bind(&input); err != nil {
		h.logger.Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if err := c.Validate(input); err != nil {
		h.logger.Error(err.Error())
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}

	authenticatedUserID := c.Get("user").(uuid.UUID)

	err := h.service.Forms.Create(sqlc.CreateFormParams{
		UserID:      authenticatedUserID,
		Title:       input.Title,
		Description: input.Description,
		Fields:      input.Fields,
		Published:   input.Published,
	})

	if err != nil {
		h.logger.Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusCreated)
}
