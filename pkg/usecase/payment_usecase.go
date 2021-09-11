package usecase

import (
	error2 "evermos_technical_test/error"
	"evermos_technical_test/pkg/domain"
	"evermos_technical_test/pkg/dto/payment"
	"evermos_technical_test/pkg/repository"
	"github.com/mitchellh/mapstructure"
	"strconv"
)

type PaymentUseCase struct {
	productRepo *repository.ProductRepository
	orderRepo   *repository.OrderRepository
	transRepo   *repository.TransactionRepository
}

func NewPaymentUseCase(
	productRepo *repository.ProductRepository,
	orderRepo *repository.OrderRepository,
	transRepo *repository.TransactionRepository) TransactionUseCase {
	return &PaymentUseCase{
		productRepo: productRepo,
		orderRepo:   orderRepo,
		transRepo:   transRepo,
	}
}

func (p PaymentUseCase) Process(data interface{}) (interface{}, error) {

	request := data.(*payment.PaymentRequest)
	var transaction *domain.Transaction
	var response *payment.PaymentResponse

	// get order
	order, err := p.orderRepo.GetOrderById(request.OrderId)
	if err != nil {
		return nil, err
	}
	//check order exist
	if order != nil {
		var initialTransaction = domain.Transaction{}
		mapstructure.Decode(request, &initialTransaction)
		initialTransaction.TransactionNumber = "INV/" + strconv.Itoa(int(initialTransaction.OrderId))
		initialTransaction.Status = "PENDING"

		for _, orderItem := range order.OrderItems {
			// get product
			product, _ := p.productRepo.GetProductById(orderItem.ItemID)

			// check available
			if product.AvailableStock < orderItem.Quantity {
				return nil, nil
			}

			product.AvailableStock -= orderItem.Quantity

			// update stock product
			go p.productRepo.UpdateProduct(product.ID, product)

		}

		// create transaction
		transaction, _ = p.transRepo.CreateTransaction(&initialTransaction)

		// check transaction success or failed
		if transaction.Amount != order.TotalAmount {
			transaction.Status = "FAILED"
			for _, orderItem := range order.OrderItems {

				// get product
				productUpdate, _ := p.productRepo.GetProductById(orderItem.ItemID)

				productUpdate.AvailableStock += orderItem.Quantity

				// update stock product
				go p.productRepo.UpdateProduct(productUpdate.ID, productUpdate)
			}

			transaction, _ := p.transRepo.UpdateTransaction(transaction.ID, transaction)
			mapstructure.Decode(transaction, &response)
			return response, nil
		} else {
			transaction.Status = "SUCCESS"
			for _, orderItem := range order.OrderItems {

				// get product
				product, _ := p.productRepo.GetProductById(orderItem.ItemID)

				product.StockSoldOut += orderItem.Quantity

				// update stock product
				go p.productRepo.UpdateProduct(product.ID, product)
			}
			transaction, _ := p.transRepo.UpdateTransaction(transaction.ID, transaction)
			mapstructure.Decode(transaction, &response)
			return response, nil
		}
	} else {
		return nil, error2.ErrBadParamInput
	}
}
