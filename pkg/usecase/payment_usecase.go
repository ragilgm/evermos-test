package usecase

import (
	error2 "evermos_technical_test/error"
	"evermos_technical_test/pkg/domain"
	"evermos_technical_test/pkg/dto/checkout"
	"evermos_technical_test/pkg/dto/payment"
	"evermos_technical_test/pkg/repository"
	"github.com/mitchellh/mapstructure"
	"time"
)

type PaymentUseCase struct {
	productRepo *repository.ProductRepository
	orderRepo   *repository.OrderRepository
	transRepo   *repository.TransactionRepository
}

func NewPaymentUseCase(
	productRepo *repository.ProductRepository,
	orderRepo *repository.OrderRepository,
	transRepo *repository.TransactionRepository) *PaymentUseCase {
	return &PaymentUseCase{
		productRepo: productRepo,
		orderRepo:   orderRepo,
		transRepo:   transRepo,
	}
}

func (p PaymentUseCase) Process(data interface{}) (interface{}, error) {

	request := data.(*payment.PaymentRequest)
	var transaction *domain.Transaction
	var response *checkout.CheckoutResponse

	// check transaction exist
	transaction, _ = p.transRepo.GetTransactionByToken(request.Token)
	if transaction == nil {
		return nil, error2.ErrBadRequest
	}

	// check status transaction
	if transaction.Status != "PENDING" {
		return nil, error2.ErrConflict
	}

	// check exipred token
	if transaction.ExpiredAt.After(time.Now()) {
		return nil, error2.ErrConflict
	}

	// get order
	order, err := p.orderRepo.GetOrderById(transaction.OrderId)
	if err != nil {
		return nil, err
	}

	//check order exist
	if order != nil {

		for _, orderItem := range order.OrderItems {
			// get product
			product, _ := p.productRepo.GetProductById(orderItem.ItemID)

			product.StockOnHold -= orderItem.Quantity
			product.StockSoldOut += orderItem.Quantity

			// update stock product
			go p.productRepo.UpdateProduct(product.ID, product)

		}

		// update status transaction
		transaction.Status = "PAID"

		transaction, _ = p.transRepo.UpdateTransaction(transaction.ID, transaction)

		// map model to response
		mapstructure.Decode(transaction, &response)

		return response, nil

	} else {
		return nil, error2.ErrBadParamInput
	}
}
