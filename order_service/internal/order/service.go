package order

import "context"

type OrderService struct {
	repo IOrderRepo
}

func NewService(repo IOrderRepo) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

type IOrderRepo interface {
	Create(ctx context.Context, order Order) (uint64, error)
	GetByID(ctx context.Context, ID uint64) (Order, error)
	GetByUserID(ctx context.Context, userID uint64) ([]Order, error)
	GetByUserIDAndStatus(ctx context.Context, userID uint64, status string) ([]Order, error)
	ChangeOrderStatus(ctx context.Context, ID uint64, status string) (uint64, error)
}

func (os *OrderService) Create(ctx context.Context, dto CreateOrderDTO) (uint64, error) {
	order := createOrderDTOToOrder(dto)
	return os.repo.Create(ctx, order)
}

func (os *OrderService) GetByID(ctx context.Context, ID uint64) (Order, error) {
	return os.repo.GetByID(ctx, ID)
}

func (os *OrderService) GetByUserID(ctx context.Context, userID uint64) ([]Order, error) {
	return os.repo.GetByUserID(ctx, userID)
}

func (os *OrderService) GetByUserIDAndStatus(ctx context.Context, userID uint64, status string) ([]Order, error) {
	return os.repo.GetByUserIDAndStatus(ctx, userID, status)
}

func (os *OrderService) ChangeOrderStatus(ctx context.Context, ID uint64, status string) (uint64, error) {
	if status == "завершен" {
		// TODO: send msg
	}

	return os.repo.ChangeOrderStatus(ctx, ID, status)
}
