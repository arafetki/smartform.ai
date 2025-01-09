package middleware

import (
	"github.com/arafetki/smartform.ai/backend/internal/config"
	"github.com/arafetki/smartform.ai/backend/internal/logging"
	"github.com/arafetki/smartform.ai/backend/internal/service"
)

type Middleware struct {
	logger  *logging.Logger
	cfg     config.Config
	service *service.Service
}

func New(logger *logging.Logger, cfg config.Config, svc *service.Service) *Middleware {
	return &Middleware{
		logger:  logger,
		cfg:     cfg,
		service: svc,
	}
}
