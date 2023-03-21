package postgres

import (
	"context"
	"fmt"

	"github.com/Levap123/order_service/internal/order"
	"github.com/jmoiron/sqlx"
)

type OrderRepo struct {
	DB *sqlx.DB
}

func NewOrderRepoPostgres(db *sqlx.DB) OrderRepo {
	return OrderRepo{
		DB: db,
	}
}

const orderTable = "orders"

func (or *OrderRepo) Create(ctx context.Context, order order.Order) (uint64, error) {
	tx, err := or.DB.BeginTxx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("order repo - create - %w", err)
	}
	defer tx.Rollback()

	query := fmt.Sprintf("INSERT INTO %s (user_id, book_id, added_at) VALUES (:user_id, :book_id, :added_at, :status) RETURNING id", orderTable)

	res, err := tx.NamedExecContext(ctx, query, order)
	if err != nil {
		return 0, fmt.Errorf("order repo - create - %w", err)
	}

	ID, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("order repo - create - %w", err)
	}

	return uint64(ID), tx.Commit()
}

func (or *OrderRepo) GetByID(ctx context.Context, ID uint64) (order.Order, error) {
	tx, err := or.DB.BeginTxx(ctx, nil)
	if err != nil {
		return order.Order{}, fmt.Errorf("order repo - get by id - %w", err)
	}
	defer tx.Rollback()

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", orderTable)

	var entity order.Order
	if err := tx.GetContext(ctx, &entity, query, ID); err != nil {
		return order.Order{}, fmt.Errorf("order repo - get by id - %w", err)
	}

	return entity, tx.Commit()
}

func (or *OrderRepo) GetByUserID(ctx context.Context, userID uint64) ([]order.Order, error) {
	tx, err := or.DB.BeginTxx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("order repo - get by user id - %w", err)
	}
	defer tx.Rollback()

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", orderTable)

	var orders []order.Order
	if err := tx.SelectContext(ctx, &orders, query, userID); err != nil {
		return nil, fmt.Errorf("order repo - get by user id - %w", err)
	}

	return orders, tx.Commit()
}

func (or *OrderRepo) GetByUserIDAndStatus(ctx context.Context, userID uint64, status string) ([]order.Order, error) {
	tx, err := or.DB.BeginTxx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("order repo - get by user id - %w", err)
	}
	defer tx.Rollback()

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1 and status = $2", orderTable)

	var orders []order.Order
	if err := tx.SelectContext(ctx, &orders, query, userID, status); err != nil {
		return nil, fmt.Errorf("order repo - get by user id - %w", err)
	}

	return orders, tx.Commit()
}

func (or *OrderRepo) ChangeOrderStatus(ctx context.Context, ID uint64, status string) (uint64, error) {
	tx, err := or.DB.BeginTxx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("order repo - change order status - %w", err)
	}
	defer tx.Rollback()

	query := fmt.Sprintf("UPDATE %s SET status = $1 WHERE id = $2", orderTable)

	if _, err := tx.ExecContext(ctx, query, status, ID); err != nil {
		return 0, fmt.Errorf("order repo - change order status - %w", err)
	}

	return ID, tx.Commit()
}
