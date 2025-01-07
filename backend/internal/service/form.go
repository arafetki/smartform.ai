package service

import (
	"errors"

	"github.com/arafetki/smartform.ai/backend/internal/db/sqlc"
	"github.com/google/uuid"
)

type formService struct {
	q *sqlc.Queries
}

var (
	ErrFormNotFound = errors.New("form not found")
)

func (s *formService) Create(params sqlc.CreateFormParams) error {
	return nil
}

func (s *formService) GetOne(id uuid.UUID) (*sqlc.Form, error) {
	return nil, nil
}
func (s *formService) GetAllForUser(userID uuid.UUID) ([]sqlc.Form, error) {
	return nil, nil
}
func (s *formService) Update(params sqlc.UpdateFormParams) error {
	return nil
}
func (s *formService) Delete(id uuid.UUID) error {
	return nil
}
