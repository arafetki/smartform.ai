package service

import (
	"errors"
	"time"

	"github.com/arafetki/smartform.ai/backend/internal/db/sqlc"
	"github.com/arafetki/smartform.ai/backend/internal/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type userService struct {
	q *sqlc.Queries
}

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrDuplicateUser = errors.New("user already exists")
)

func (s *userService) Create(params sqlc.CreateUserParams) error {
	ctx, cancel := utils.ContextWithTimeout(5 * time.Second)
	defer cancel()

	err := s.q.CreateUser(ctx, params)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.ConstraintName {
			case "users_pkey":
				return ErrDuplicateUser
			default:
				return pgErr
			}
		}
		return err
	}

	return nil
}

func (s *userService) GetOne(id uuid.UUID) (*sqlc.User, error) {
	ctx, cancel := utils.ContextWithTimeout(5 * time.Second)
	defer cancel()
	user, err := s.q.GetUser(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (s *userService) Update(params sqlc.UpdateUserParams) error {
	ctx, cancel := utils.ContextWithTimeout(5 * time.Second)
	defer cancel()

	rowsAffected, err := s.q.UpdateUser(ctx, params)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrUserNotFound
	}
	return nil
}

func (s *userService) Delete(id uuid.UUID) error {
	ctx, cancel := utils.ContextWithTimeout(5 * time.Second)
	defer cancel()
	rowsAffected, err := s.q.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrUserNotFound
	}
	return nil
}
