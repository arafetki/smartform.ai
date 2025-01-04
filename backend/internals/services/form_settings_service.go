package services

import (
	"time"

	"github.com/arafetki/smartform.ai/backend/internals/repository/sqlc"
	"github.com/arafetki/smartform.ai/backend/internals/utils"
)

type FormSettingsService struct {
	q *sqlc.Queries
}

func NewFormSettingsService(q *sqlc.Queries) *FormSettingsService {
	return &FormSettingsService{q}
}

func (s *FormSettingsService) ListAllSettings() ([]sqlc.FormSettings, error) {
	ctx, cancel := utils.ContextWithTimeout(5 * time.Second)
	defer cancel()
	return s.q.ListSettings(ctx)
}
