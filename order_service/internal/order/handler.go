package order

import (
	"context"

	"github.com/Levap123/order_service/proto"
	"github.com/sirupsen/logrus"
)

type OrderHandler struct {
	repo IOrderRepo
	proto.UnimplementedOrdersServer
	logger *logrus.Logger
}


func (h *OrderHandler) Create(ctx context.Context, req *proto.CreateOrderRequest) (*proto.CreateOrderResponse, error) {
	dto := fromReqToCreateDTO(req)

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
