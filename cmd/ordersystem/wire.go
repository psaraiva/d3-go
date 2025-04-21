//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"psaraiva/d3/internal/entity"
	"psaraiva/d3/internal/event"
	"psaraiva/d3/internal/infra/database"
	"psaraiva/d3/internal/infra/web"
	"psaraiva/d3/internal/usecase"
	"psaraiva/d3/pkg/events"

	"github.com/google/wire"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	event.NewOrderUpdated,
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventInterface), new(*event.OrderUpdated)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

var setOrderUpdatedEvent = wire.NewSet(
	event.NewOrderUpdated,
	wire.Bind(new(events.EventInterface), new(*event.OrderUpdated)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewUpdateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.UpdateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderUpdatedEvent,
		usecase.NewUpdateOrderUseCase,
	)
	return &usecase.UpdateOrderUseCase{}
}

var setOrderEvents = wire.NewSet(
	event.NewOrderCreated,
	event.NewOrderUpdated,
	wire.Struct(new(event.OrderEvents), "OrderCreated", "OrderUpdated"),
)

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderEvents,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}
