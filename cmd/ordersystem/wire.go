//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/VictorH97/devfullcycle/goexpert/CleanArch/internal/entity"
	"github.com/VictorH97/devfullcycle/goexpert/CleanArch/internal/event"
	"github.com/VictorH97/devfullcycle/goexpert/CleanArch/internal/infra/database"
	"github.com/VictorH97/devfullcycle/goexpert/CleanArch/internal/infra/web"
	"github.com/VictorH97/devfullcycle/goexpert/CleanArch/internal/usecase"
	"github.com/VictorH97/devfullcycle/goexpert/CleanArch/pkg/events"
	"github.com/google/wire"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

var setOrdersListedEvent = wire.NewSet(
	event.NewOrdersListed,
	wire.Bind(new(events.EventInterface), new(*event.OrdersListed)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}

func NewListOrdersUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.ListOrdersUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrdersListedEvent,
		usecase.NewListOrdersUseCase,
	)
	return &usecase.ListOrdersUseCase{}
}
