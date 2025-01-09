// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlc

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateForm(ctx context.Context, arg CreateFormParams) error
	CreateUser(ctx context.Context, arg CreateUserParams) error
	DeleteForm(ctx context.Context, id uuid.UUID) error
	DeleteUser(ctx context.Context, id uuid.UUID) (int64, error)
	GetForm(ctx context.Context, id uuid.UUID) (Form, error)
	GetUser(ctx context.Context, id uuid.UUID) (User, error)
	ListFormsForUser(ctx context.Context, userID uuid.UUID) ([]ListFormsForUserRow, error)
	UpdateForm(ctx context.Context, arg UpdateFormParams) (int64, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (int64, error)
}

var _ Querier = (*Queries)(nil)
