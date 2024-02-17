package usecase

import (
	"github.com/VictorH97/devfullcycle/goexpert/CleanArch/internal/entity"
	"github.com/VictorH97/devfullcycle/goexpert/CleanArch/pkg/events"
)

type ListOrdersOutput struct {
	Orders []OrderOutputDTO
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrdersListed    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrdersListed events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
		OrdersListed:    OrdersListed,
		EventDispatcher: EventDispatcher,
	}
}

func (r *ListOrdersUseCase) Execute() (ListOrdersOutput, error) {
	orders, err := r.OrderRepository.FindAll()
	if err != nil {
		return ListOrdersOutput{}, err
	}

	var listOrders []OrderOutputDTO

	for _, order := range orders {
		listOrders = append(listOrders, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	r.OrdersListed.SetPayload(listOrders)
	r.EventDispatcher.Dispatch(r.OrdersListed)

	return ListOrdersOutput{
		Orders: listOrders,
	}, nil
}
