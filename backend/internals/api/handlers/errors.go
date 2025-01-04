package handlers

import (
	"log/slog"
	"net/http"
	"runtime/debug"

	"github.com/arafetki/smartform.ai/backend/internals/logging"
	"github.com/labstack/echo/v4"
)

func logServerError(r *http.Request, err error) {
	var (
		message string = err.Error()
		method  string = r.Method
		url     string = r.URL.String()
		trace   string = string(debug.Stack())
	)
	requestAttrs := slog.Group("request", "method", method, "url", url)
	logging.Logger().Error(message, requestAttrs, "trace", trace)
}

func (h *Handler) internalServerErrorResponse(c echo.Context, err error) error {
	message := "The server encountered a problem and could not process your request."
	logServerError(c.Request(), err)
	return echo.NewHTTPError(http.StatusInternalServerError, message)
}

func (h *Handler) notFoundErrorResponse() error {
	message := "The requested resource could not be found."
	return echo.NewHTTPError(http.StatusNotFound, message)
}

func (h *Handler) badRequestErrorResponse(c echo.Context, err error) error {
	message := "The request could not be understood or was missing required parameters."
	logServerError(c.Request(), err)
	return echo.NewHTTPError(http.StatusBadRequest, message)
}
