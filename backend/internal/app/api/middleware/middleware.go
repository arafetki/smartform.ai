package middleware

import (
	"github.com/arafetki/smartform.ai/backend/internal/config"
	"github.com/arafetki/smartform.ai/backend/internal/logging"
)

type Middleware struct {
	logger *logging.Logger
	cfg    config.Config
}

func New(logger *logging.Logger, cfg config.Config) *Middleware {
	return &Middleware{
		logger: logger,
		cfg:    cfg,
	}
}
