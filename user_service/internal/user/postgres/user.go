package postgres

import (
	"context"
	"fmt"
	"user_service/internal/user"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	DB *sqlx.DB
}

const userTable = "users"

func (ur *UserRepo) Create(ctx context.Context, user *user.User) (int, error) {
	tx, err := ur.DB.BeginTxx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("user repo - create - start tx - %w", err)
	}
	defer func() { tx.Rollback() }()
	query := fmt.Sprintf("INSERT INTO %s(email, username, password) VALUES ($1, $2, $3) RETURNING id", userTable)
	var userID int
	if err := tx.Get(&userID, query, user.Email, user.Username, user.Password); err != nil {
		return 0, fmt.Errorf("user repo - create - insert - %w", err)
	}
	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("user repo - create - commit tx - %w", err)
	}
	return userID, nil
}
