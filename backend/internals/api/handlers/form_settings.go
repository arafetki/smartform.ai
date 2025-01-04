package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) FetchAllSettings(c echo.Context) error {
	settings, err := h.FormSettingsService.ListAllSettings()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, echo.Map{"data": settings})
}
