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

type UsersService struct {
	q *sqlc.Queries
}

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrDuplicateEmail       = errors.New("email already exists")
	ErrDuplicatePhoneNumber = errors.New("phone number already exists")
)

func NewUsersService(q *sqlc.Queries) *UsersService {
	return &UsersService{q}
}

func (s *UsersService) CreateUser(params sqlc.CreateUserParams) error {
	ctx, cancel := utils.ContextWithTimeout(5 * time.Second)
	defer cancel()

	err := s.q.CreateUser(ctx, params)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.ConstraintName {
			case "users_email_key":
				return ErrDuplicateEmail
			case "users_phone_number_key":
				return ErrDuplicatePhoneNumber
			default:
				return pgErr
			}
		}
		return err
	}

	return nil
}

func (s *UsersService) GetUserByID(id uuid.UUID) (*sqlc.User, error) {
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

func (s *UsersService) ListAllUsers() ([]sqlc.User, error) {
	ctx, cancel := utils.ContextWithTimeout(5 * time.Second)
	defer cancel()
	return s.q.ListUsers(ctx)
}

func (s *UsersService) UpdateUser(params sqlc.UpdateUserParams) error {
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

func (s *UsersService) DeleteUser(id uuid.UUID) error {
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

func (s *UsersService) CountUsers() (int64, error) {
	ctx, cancel := utils.ContextWithTimeout(5 * time.Second)
	defer cancel()

	return s.q.CountUsers(ctx)
}
