package services

import (
	"errors"
	"time"

	"github.com/arafetki/smartform.ai/backend/internals/repository/sqlc"
	"github.com/arafetki/smartform.ai/backend/internals/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type FormsService struct {
	q *sqlc.Queries
}

var (
	ErrFormNotFound    = errors.New("form not found")
	ErrNoFormsDeleted  = errors.New("no forms deleted")
	ErrActionForbidden = errors.New("action forbidden")
)

func NewFormsService(q *sqlc.Queries) *FormsService {
	return &FormsService{q}
}

func (s *FormsService) CreateForm(params sqlc.CreateFormParams) error {
	ctx, cancel := utils.ContextWithTimeout(5 * time.Second)
	defer cancel()
	return s.q.CreateForm(ctx, params)
}

func (s *FormsService) DeleteForms(params sqlc.DeleteFormsByOwnerParams) error {
	ctx, cancel := utils.ContextWithTimeout(5 * time.Second)
	defer cancel()

	rowsAffected, err := s.q.DeleteFormsByOwner(ctx, params)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrNoFormsDeleted
	}
	return nil
}

func (s *FormsService) GetFormWithSettings(id uuid.UUID) (*sqlc.GetFormWithSettingsRow, error) {
	ctx, cancel := utils.ContextWithTimeout(5 * time.Second)
	defer cancel()
	data, err := s.q.GetFormWithSettings(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrFormNotFound
		}
		return nil, err
	}
	return &data, nil
}

func (s *FormsService) ListFormsForUser(userID uuid.UUID) ([]sqlc.ListFormsForUserRow, error) {
	ctx, cancel := utils.ContextWithTimeout(5 * time.Second)
	defer cancel()

	return s.q.ListFormsForUser(ctx, userID)
}

func (s *FormsService) UpdateForm(params sqlc.UpdateFormParams) error {
	ctx, cancel := utils.ContextWithTimeout(5 * time.Second)
	defer cancel()

	rowsAffected, err := s.q.UpdateForm(ctx, params)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrFormNotFound
	}
	return nil
}
