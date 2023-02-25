package postgres

import "github.com/jmoiron/sqlx"

type UserRepo struct {
	DB *sqlx.DB
}

func (ur *UserRepo) Create()
