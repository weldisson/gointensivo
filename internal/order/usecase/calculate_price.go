package usecase

import (
	"github.com/weldisson/gointensivo/internal/order/entity"
)

// DTO(data-transfer-object) são usados para transferir dados de uma camada para outra da aplicação
type OrderInputDTO struct {
	ID    string
	Price float64
	Tax   float64
}

type OrderOutputDTO struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

type CalculateFinalPriceUseCase struct {
	/* THE BIG COUPLING,
	para não ter esse acoplamento abaixo é necessário depender de uma interface
	*/
	// OrderRepository database.OrderRepository
	OrderRepository entity.OrderRepositoryInterface // sendo essa a melhor forma.
}

func NewCalculateFinalPriceUseCase(orderRepository entity.OrderRepositoryInterface) *CalculateFinalPriceUseCase {
	return &CalculateFinalPriceUseCase{
		// OrderRepository: orderRepository, // dessa forma o orderRepository não está implementando a OrderRepositoryInterface
		OrderRepository: orderRepository, // dessa forma o orderRepository está implementando a OrderRepositoryInterface
	}
}

// essa func abaixo é um método da struct CalculateFinalPriceUseCase
func (c *CalculateFinalPriceUseCase) Execute(input OrderInputDTO) (*OrderOutputDTO, error) {
	order, err := entity.NewOrder(input.ID, input.Price, input.Tax)
	if err != nil {
		return nil, err
	}
	err = order.CalculateFinalPrice()
	if err != nil {
		return nil, err
	}
	err = c.OrderRepository.Save(order)
	if err != nil {
		return nil, err
	}

	return &OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
