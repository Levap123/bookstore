package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"user_service/internal/domain"
	"user_service/internal/user"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type UserRepo struct {
	DB *sqlx.DB
	lg *logrus.Logger
}

func NewUserRepo(DB *sqlx.DB, lg *logrus.Logger) *UserRepo {
	return &UserRepo{
		DB: DB,
		lg: lg,
	}
}

const userTable = "users"

func (ur *UserRepo) Create(ctx context.Context, user *user.User) (uint64, error) {
	tx, err := ur.DB.BeginTxx(ctx, nil)
	if err != nil {
		ur.lg.Error(err)
		return 0, fmt.Errorf("user repo create - start tx - %w", err)
	}
	defer tx.Rollback()

	query := fmt.Sprintf("INSERT INTO %s(email, username, password) VALUES ($1, $2, $3) RETURNING id", userTable)
	var userID uint64
	if err := tx.Get(&userID, query, user.Email, user.Username, user.Password); err != nil {
		ur.lg.Error(err)
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return 0, fmt.Errorf("user repo create - insert - %w", domain.ErrUnique)
		}
		return 0, fmt.Errorf("user repo create - insert - %w", err)
	}

	if err := tx.Commit(); err != nil {
		ur.lg.Error(err)
		return 0, fmt.Errorf("user repo create - commit tx - %w", err)
	}
	return userID, nil
}

func (ur *UserRepo) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	tx, err := ur.DB.BeginTxx(ctx, nil)
	if err != nil {
		ur.lg.Error(err)
		return nil, fmt.Errorf("user repo get - start tx - %w", err)
	}

	defer tx.Rollback()

	query := fmt.Sprintf("SELECT * FROM %s WHERE email = $1", userTable)
	var user user.User

	if err := tx.Get(&user, query, email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user repo get - select - %w", domain.ErrUserNotFound)
		}
		return nil, fmt.Errorf("user repo get - select - %w", err)
	}

	if err := tx.Commit(); err != nil {
		ur.lg.Error(err)
		return nil, fmt.Errorf("user repo get by email - commit tx - %w", err)
	}

	return &user, nil
}

func (ur *UserRepo) GetByID(ctx context.Context, ID uint64) (*user.User, error) {
	tx, err := ur.DB.BeginTxx(ctx, nil)
	if err != nil {
		ur.lg.Error(err)
		return nil, fmt.Errorf("user repo - get by ID - start tx - %w", err)
	}

	defer tx.Rollback()

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", userTable)
	var user user.User

	if err := tx.Get(&user, query, ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user repo - get by ID - select - %w", domain.ErrUserNotFound)
		}
		return nil, fmt.Errorf("user repo get - by ID - select - %w", err)
	}

	if err := tx.Commit(); err != nil {
		ur.lg.Error(err)
		return nil, fmt.Errorf("user repo get by id - commit tx - %w", err)
	}

	return &user, nil
}

func (ur *UserRepo) UpdateInfo(ctx context.Context, user *user.User) (int, error) {
	tx, err := ur.DB.BeginTxx(ctx, nil)
	if err != nil {
		ur.lg.Error(err)
		return 0, fmt.Errorf("user repo get - start tx - %w", err)
	}

	defer tx.Rollback()

	query := fmt.Sprintf("UPDATE %s SET username = $1, password = $2 WHERE id = $3 RETURNING id", userTable)
	var userID int

	if err := tx.Get(&userID, query, user.Username, user.Password, user.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("user repo - update user - select - %w", domain.ErrUserNotFound)
		}
	}

	if err := tx.Commit(); err != nil {
		ur.lg.Error(err)
		return 0, fmt.Errorf("user update info - commit tx - %w", err)
	}

	return userID, nil
}
