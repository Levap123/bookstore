package postgres

import (
	"fmt"

	"github.com/Levap123/order_service/internal/configs"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func InitDB(cfg *configs.Configs) (*sqlx.DB, error) {
	connURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.Postgres.Username, cfg.Postgres.Password, cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.DBName)
	return sqlx.Open("pgx", connURL)
}
