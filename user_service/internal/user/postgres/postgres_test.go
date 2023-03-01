package postgres_test

import (
	"fmt"
	"os"
	"testing"

	_ "github.com/jackc/pgx/stdlib"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func TestMain(m *testing.M) {
	code, err := run(m)
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}

func run(m *testing.M) (int, error) {
	var err error
	DB, err = sqlx.Open("pgx", "postgres://test:test@localhost:5000/testdb?sslmode=disable")
	if err != nil {
		return -1, err
	}
	if err := DB.Ping(); err != nil {
		return -1, err
	}
	if _, err := DB.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email TEXT UNIQUE NOT NULL,
		username TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	);`); err != nil {
		return -1, err
	}
	defer DB.Exec("DROP TABLE users")

	defer DB.Close()

	return m.Run(), nil
}
