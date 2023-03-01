package postgres

import (
	"fmt"
	"os"
	"testing"

	_ "github.com/jackc/pgx/stdlib"

	"github.com/jmoiron/sqlx"
)

func TestMain(m *testing.M) {
	code, err := run(m)
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}

func run(m *testing.M) (int, error) {
	DB, err := sqlx.Open("pgx", "postgres://test:test@localhost:5000/testdb")
	if err != nil {
		return -1, err
	}

	defer func() {
		for _, t := range []string{"users"} {
			_, _ = DB.Exec(fmt.Sprintf("DELETE FROM %s", t))
		}

		DB.Close()
	}()
	return m.Run(), nil
}
