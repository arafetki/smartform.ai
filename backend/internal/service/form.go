package service

import (
	"errors"
	"time"

	"github.com/arafetki/smartform.ai/backend/internal/db/sqlc"
	"github.com/arafetki/smartform.ai/backend/internal/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type formService struct {
	q *sqlc.Queries
}

var (
	ErrFormNotFound = errors.New("form not found")
	ErrUnauthorized = errors.New("unauthorized")
)

func (s *formService) Create(params sqlc.CreateFormParams) error {
	ctx, cancel := utils.ContextWithTimeout(3 * time.Second)
	defer cancel()

	return s.q.CreateForm(ctx, params)
}

func (s *formService) GetOne(id uuid.UUID) (*sqlc.Form, error) {
	ctx, cancel := utils.ContextWithTimeout(3 * time.Second)
	defer cancel()

	form, err := s.q.GetForm(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrFormNotFound
		}
		return nil, err
	}
	return &form, nil
}
func (s *formService) GetAllForUser(userID uuid.UUID) ([]sqlc.ListFormsForUserRow, error) {
	ctx, cancel := utils.ContextWithTimeout(3 * time.Second)
	defer cancel()

	return s.q.ListFormsForUser(ctx, userID)

}
func (s *formService) Update(params sqlc.UpdateFormParams) error {
	ctx, cancel := utils.ContextWithTimeout(3 * time.Second)
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
func (s *formService) Delete(id uuid.UUID, ownerId uuid.UUID) error {

	ctx, cancel := utils.ContextWithTimeout(3 * time.Second)
	defer cancel()

	form, err := s.GetOne(id)
	if err != nil {
		return err
	}
	if form.UserID != ownerId {
		return ErrUnauthorized
	}

	return s.q.DeleteForm(ctx, id)
}
