package handler

import (
	"github.com/arafetki/smartform.ai/backend/internal/config"
	"github.com/arafetki/smartform.ai/backend/internal/logging"
	"github.com/arafetki/smartform.ai/backend/internal/service"
)

type Handler struct {
	logger  *logging.Logger
	cfg     config.Config
	service *service.Service
}

func New(logger *logging.Logger, cfg config.Config, svc *service.Service) *Handler {
	return &Handler{
		logger:  logger,
		cfg:     cfg,
		service: svc,
	}
}
