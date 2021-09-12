package usecase

import (
	error2 "evermos_technical_test/error"
	"evermos_technical_test/pkg/domain"
	"evermos_technical_test/pkg/dto/checkout"
	"evermos_technical_test/pkg/dto/payment"
	"evermos_technical_test/pkg/repository"
	"github.com/mitchellh/mapstructure"
	"sync"
	"time"
)

type PaymentUseCase struct {
	mutex       sync.Mutex
	productRepo *repository.ProductRepository
	orderRepo   *repository.OrderRepository
	transRepo   *repository.TransactionRepository
}

func NewPaymentUseCase(
	mutex sync.Mutex,
	productRepo *repository.ProductRepository,
	orderRepo *repository.OrderRepository,
	transRepo *repository.TransactionRepository) *PaymentUseCase {
	return &PaymentUseCase{
		mutex:       mutex,
		productRepo: productRepo,
		orderRepo:   orderRepo,
		transRepo:   transRepo,
	}
}

const (
	StatusPaid    = "PAID"
	StatusPending = "PENDING"
)

func (p PaymentUseCase) Process(data interface{}) (interface{}, error) {

	request := data.(*payment.PaymentRequest)
	var transaction *domain.Transaction
	var response *checkout.CheckoutResponse

	// check transaction exist
	transaction, err := p.transRepo.GetTransactionByToken(request.Token)

	if err != nil {
		return nil, err
	}

	// check transaction exist
	if transaction == nil {
		return nil, error2.ErrNotFound
	}

	// check status transaction
	if transaction.Status != StatusPending {
		return nil, error2.ErrConflict
	}

	// check exipred token
	if transaction.ExpiredAt.After(time.Now()) {
		return nil, error2.TokenExpired
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

			go func() {
				// update stock product
				p.mutex.Lock()
				_, err := p.productRepo.UpdateProduct(product.ID, product)
				p.mutex.Unlock()
				if err != nil {
				}
			}()

		}

		// update status transaction
		transaction.Status = StatusPaid

		transaction, _ = p.transRepo.UpdateTransaction(transaction.ID, transaction)

		// map model to response
		err := mapstructure.Decode(transaction, &response)
		if err != nil {
			return nil, err
		}

		return response, nil

	} else {
		return nil, error2.ErrBadParamInput
	}
}
