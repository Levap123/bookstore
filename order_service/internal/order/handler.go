package order

import (
	"context"

	"github.com/Levap123/order_service/proto"
	"github.com/sirupsen/logrus"
)

type OrderHandler struct {
	service IOrderService
	proto.UnimplementedOrdersServer
	logger *logrus.Logger
}

type IOrderService interface {
	Create(ctx context.Context, dto CreateOrderDTO) (uint64, error)
	GetByID(ctx context.Context, ID uint64) (Order, error)
	GetByUserID(ctx context.Context, userID uint64) ([]Order, error)
	GetByUserIDAndStatus(ctx context.Context, userID uint64, status string) ([]Order, error)
	ChangeOrderStatus(ctx context.Context, ID uint64, status string) (uint64, error)
}

func (h *OrderHandler) Create(ctx context.Context, req *proto.CreateOrderRequest) (*proto.CreateOrderResponse, error) {
	dto := fromReqToCreateDTO(req)

	resp, err := h.service.Create(ctx, dto)
	if err != nil {
		h.logger.Error("error in creating order: %v", err)

		return nil, err
	}

	return &proto.CreateOrderResponse{
		Id: resp,
	}, nil
}

func (h *OrderHandler) GetByUserID(ctx context.Context, req *proto.GetOrderByUserIDRequest) (*proto.OrderArray, error) {
	resp, err := h.service.GetByUserID(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	respArr := make([]*proto.Order, 0, len(resp))

	for _, order := range resp {
		respArr = append(respArr, fromOrderToResp(order))
	}

	return &proto.OrderArray{
		Oo: respArr,
	}, nil
}

func (h *OrderHandler) GetByID(ctx context.Context, req *proto.GetOrderByIDRequest) (*proto.Order, error) {
	order, err := h.service.GetByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return fromOrderToResp(order), nil
}

func (h *OrderHandler) ChangeStatus(ctx context.Context, req *proto.ChangeStatusRequest) (*proto.CreateOrderResponse, error) {
	resp, err := h.service.ChangeOrderStatus(ctx, req.Id, req.Status)
	if err != nil {
		return nil, err
	}

	return &proto.CreateOrderResponse{
		Id: resp,
	}, nil
}

func (h *OrderHandler) GetByUserIDAndStatus(ctx context.Context, req *proto.GetOrderByUserIDAndStatusRequest) (*proto.OrderArray, error) {
	resp, err := h.service.GetByUserIDAndStatus(ctx, req.UserId, req.Status)
	if err != nil {
		return nil, err
	}

	respArr := make([]*proto.Order, 0, len(resp))

	for _, order := range resp {
		respArr = append(respArr, fromOrderToResp(order))
	}

	return &proto.OrderArray{
		Oo: respArr,
	}, nil
}

/*
type OrdersServer interface {
	mustEmbedUnimplementedOrdersServer()
}
*/
