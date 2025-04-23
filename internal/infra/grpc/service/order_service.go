package service

import (
	"context"

	"psaraiva/d3/internal/infra/grpc/pb"
	"psaraiva/d3/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	UpdateOrderUseCase usecase.UpdateOrderUseCase
	ListOrderUseCase   usecase.ListOrderUseCase
}

func NewOrderService(
	createOrderUseCase usecase.CreateOrderUseCase,
	updateOrderUseCase usecase.UpdateOrderUseCase,
	listOrderUseCase usecase.ListOrderUseCase,
) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		UpdateOrderUseCase: updateOrderUseCase,
		ListOrderUseCase:   listOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderCreateInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}

	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}

	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) UpdateOrder(ctx context.Context, in *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	dto := usecase.OrderUpdateInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}

	output, err := s.UpdateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrder(ctx context.Context, in *pb.Empty) (*pb.ListOrderResponse, error) {
	output, err := s.ListOrderUseCase.Execute()
	if err != nil {
		return nil, err
	}

	ListOrderItemResponses := make([]*pb.ListOrderItemResponse, len(output.List))
	for i, item := range output.List {
		ListOrderItemResponses[i] = &pb.ListOrderItemResponse{
			Id:         item.ID,
			Price:      float32(item.Price),
			Tax:        float32(item.Tax),
			FinalPrice: float32(item.FinalPrice),
		}
	}

	return &pb.ListOrderResponse{
		List: ListOrderItemResponses,
	}, nil
}
