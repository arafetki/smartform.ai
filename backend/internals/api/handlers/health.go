package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (h *Handler) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"timestamp": time.Now().UnixNano()})
}
