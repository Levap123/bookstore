package order

import "time"

type CreateOrderDTO struct {
	BookID string
	UserID int
}

type Order struct {
	ID      int64     `db:"id"`
	BookID  string    `db:"book_id"`
	UserID  int       `db:"user_id"`
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
