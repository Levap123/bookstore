package postgres

import (
	"context"
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
	defer func() { err = tx.Rollback() }()
	query := fmt.Sprintf("INSERT INTO %s(email, username, password) VALUES ($1, $2, $3) RETURNING id", userTable)
	var userID uint64
	if err := tx.Get(&userID, query, user.Email, user.Username, user.Password); err != nil {
		ur.lg.Error(err)
		if strings.Contains(err.Error(), "duplicate key value") {
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

// func (ur *UserRepo)
