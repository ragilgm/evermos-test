package usecase

import (
	error2 "evermos_technical_test/error"
	"evermos_technical_test/pkg/domain"
	"evermos_technical_test/pkg/dto/order"
	"evermos_technical_test/pkg/repository"
	"github.com/mitchellh/mapstructure"
)

type OrderUseCase struct {
	productRepo *repository.ProductRepository
	orderRepo   *repository.OrderRepository
	transRepo   *repository.TransactionRepository
}

func NewOrderUseCase(productRepo *repository.ProductRepository, orderRepo *repository.OrderRepository) *OrderUseCase {
	return &OrderUseCase{
		productRepo: productRepo,
		orderRepo:   orderRepo,
	}
}

func (o OrderUseCase) CreateOrder(req *order.OrderRequest) (res *order.OrderResponse, err error) {
	var order = domain.Order{}
	order.Status = "STORED"
	for _, items := range req.OrderItems {
		var product, err = o.productRepo.GetProductById(items.ItemID)
		if err != nil {
			return nil, error2.ErrNotFound
		}

		if items.Quantity > product.AvailableStock {
			return nil, error2.ErrBadParamInput
		}

		item := domain.Item{ID: product.ID, ItemName: product.ProductName, Amount: product.Price}
		var orderItem = domain.OrderItem{OrderID: order.ID, ItemID: product.ID, Item: item, Quantity: items.Quantity}
		orderItem.Total = orderItem.Quantity * orderItem.Item.Amount
		order.OrderItems = append(order.OrderItems, orderItem)
		order.TotalAmount += orderItem.Total
	}

	result, err := o.orderRepo.CreateOrder(order)
	if err != nil {
		return nil, error2.ErrBadParamInput
	}

	mapstructure.Decode(result, &res)

	return res, nil
}
