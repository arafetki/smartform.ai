package service

import (
	"github.com/arafetki/smartform.ai/backend/internal/db/sqlc"
	"github.com/google/uuid"
)

type Service struct {
	Users interface {
		Create(params sqlc.CreateUserParams) error
		GetOne(id uuid.UUID) (*sqlc.User, error)
		Update(params sqlc.UpdateUserParams) error
		Delete(id uuid.UUID) error
	}
	Forms interface {
		Create(params sqlc.CreateFormParams) error
		GetOne(id uuid.UUID) (*sqlc.Form, error)
		GetAllForUser(userID uuid.UUID) ([]sqlc.Form, error)
		Update(params sqlc.UpdateFormParams) error
		Delete(id uuid.UUID) error
	}
}

func New(q *sqlc.Queries) *Service {
	return &Service{
		Users: &userService{q},
		Forms: &formService{q},
	}
}
