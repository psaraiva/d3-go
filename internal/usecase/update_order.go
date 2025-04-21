package usecase

import (
	"psaraiva/d3/internal/entity"
	"psaraiva/d3/pkg/events"
)

type OrderUpdateInputDTO struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

type OrderUpdateOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type UpdateOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderUpdated    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewUpdateOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrderUpdated events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *UpdateOrderUseCase {
	return &UpdateOrderUseCase{
		OrderRepository: OrderRepository,
		OrderUpdated:    OrderUpdated,
		EventDispatcher: EventDispatcher,
	}
}

func (c *UpdateOrderUseCase) Execute(input OrderUpdateInputDTO) (OrderUpdateOutputDTO, error) {
	order := entity.Order{
		ID:    input.ID,
		Price: input.Price,
		Tax:   input.Tax,
	}

	order.CalculateFinalPrice()
	if err := c.OrderRepository.Update(&order); err != nil {
		return OrderUpdateOutputDTO{}, err
	}

	dto := OrderUpdateOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}

	c.OrderUpdated.SetPayload(dto)
	c.EventDispatcher.Dispatch(c.OrderUpdated)

	return dto, nil
}
