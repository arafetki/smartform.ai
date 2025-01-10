package service

import (
	"context"

	"github.com/arafetki/smartform.ai/backend/internal/db/sqlc"
	"github.com/google/uuid"
)

type Service struct {
	Users interface {
		Create(ctx context.Context, params sqlc.CreateUserParams) error
		GetOne(ctx context.Context, id string) (*sqlc.User, error)
		Update(ctx context.Context, params sqlc.UpdateUserParams) error
		Delete(ctx context.Context, id string) error
	}
	Forms interface {
		Create(ctx context.Context, params sqlc.CreateFormParams) error
		GetOne(ctx context.Context, id uuid.UUID) (*sqlc.Form, error)
		GetAllForUser(ctx context.Context, userID string) ([]sqlc.Form, error)
		Update(ctx context.Context, params sqlc.UpdateFormParams) error
		Delete(ctx context.Context, Id uuid.UUID) error
	}
}

func New(q *sqlc.Queries) *Service {
	return &Service{
		Users: &userService{q},
	}
}
