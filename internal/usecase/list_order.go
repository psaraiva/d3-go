package usecase

import (
	"psaraiva/d3/internal/entity"
	"psaraiva/d3/pkg/events"
)

type OrderListItemOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type OrderListOutputDTO struct {
	List []OrderListItemOutputDTO `json:"list"`
}

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderListed     events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrderListed events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
		OrderListed:     OrderListed,
		EventDispatcher: EventDispatcher,
	}
}

func (c *ListOrderUseCase) Execute() (OrderListOutputDTO, error) {
	orders, err := c.OrderRepository.List()
	if err != nil {
		return OrderListOutputDTO{}, err
	}

	if len(*orders) == 0 {
		return OrderListOutputDTO{}, nil
	}

	listDTO := OrderListOutputDTO{}
	for _, order := range *orders {
		dto := OrderListItemOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
		listDTO.List = append(listDTO.List, dto)
	}

	c.OrderListed.SetPayload(listDTO)
	c.EventDispatcher.Dispatch(c.OrderListed)
	return listDTO, nil
}
