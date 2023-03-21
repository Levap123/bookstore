package order

import "time"

type CreateOrderDTO struct {
	BookID string
	UserID uint64
}

type Order struct {
	ID      uint64    `db:"id"`
	BookID  string    `db:"book_id"`
	UserID  uint64    `db:"user_id"`
	AddedAt time.Time `db:"added_at"`
	Status  string    `db:"status"`
}

func CreateOrderDTOToOrder(dto CreateOrderDTO) Order {
	return Order{
		BookID:  dto.BookID,
		UserID:  dto.UserID,
		Status:  "в отправке",
		AddedAt: time.Now(),
	}
}
