package web

import (
	"encoding/json"
	"net/http"

	"psaraiva/d3/internal/entity"
	"psaraiva/d3/internal/event"
	"psaraiva/d3/internal/usecase"
	"psaraiva/d3/pkg/events"
)

type WebOrderHandler struct {
	EventDispatcher events.EventDispatcherInterface
	OrderRepository entity.OrderRepositoryInterface
	OrderEvents     *event.OrderEvents
}

func NewWebOrderHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
	OrderEvents *event.OrderEvents,
) *WebOrderHandler {
	return &WebOrderHandler{
		EventDispatcher: EventDispatcher,
		OrderRepository: OrderRepository,
		OrderEvents:     OrderEvents,
	}
}

func (h *WebOrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.OrderCreateInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createOrder := usecase.NewCreateOrderUseCase(h.OrderRepository, h.OrderEvents.OrderCreated, h.EventDispatcher)
	output, err := createOrder.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *WebOrderHandler) Update(w http.ResponseWriter, r *http.Request) {
	var dto usecase.OrderUpdateInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updateOrder := usecase.NewUpdateOrderUseCase(h.OrderRepository, h.OrderEvents.OrderUpdated, h.EventDispatcher)
	output, err := updateOrder.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
