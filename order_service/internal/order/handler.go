package order

import "context"

type OrderHandler struct {
	repo IOrderRepo
}

type IOrderRepo interface {
	Create(ctx context.Context, order Order) (int64, error)
	GetByID(ctx context.Context, ID int) (Order, error)
	GetByUserID(ctx context.Context, userID int) ([]Order, error)
	GetByUserIDAndStatus(ctx context.Context, userID int, status string) ([]Order, error)
	ChangeOrderStatus(ctx context.Context, ID int, status string) (int, error)
}

/*
type OrdersServer interface {
	Create(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)
	GetByUserID(context.Context, *GetOrderByIDRequest) (*OrderArray, error)
	GetByID(context.Context, *GetOrderByIDRequest) (*Order, error)
	ChangeStatus(context.Context, *ChangeStatusRequest) (*CreateOrderResponse, error)
	GetByUserIDAndStatus(context.Context, *GetOrderByUserIDAndStatusRequest) (*OrderArray, error)
	mustEmbedUnimplementedOrdersServer()
}
*/