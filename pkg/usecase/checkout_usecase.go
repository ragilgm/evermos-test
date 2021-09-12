package usecase

import (
	error2 "evermos_technical_test/error"
	"evermos_technical_test/pkg/domain"
	"evermos_technical_test/pkg/dto/checkout"
	"evermos_technical_test/pkg/repository"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type CheckOutUseCase struct {
	productRepo *repository.ProductRepository
	orderRepo   *repository.OrderRepository
	transRepo   *repository.TransactionRepository
	mutex       sync.Mutex
}

func NewCheckOutUseCase(
	mutex sync.Mutex,
	productRepo *repository.ProductRepository,
	orderRepo *repository.OrderRepository,
	transRepo *repository.TransactionRepository) *CheckOutUseCase {
	return &CheckOutUseCase{
		mutex:       mutex,
		productRepo: productRepo,
		orderRepo:   orderRepo,
		transRepo:   transRepo,
	}
}

const (
	defaultFormatInvoice = "INV/"
)

func (c CheckOutUseCase) Process(data interface{}) (interface{}, error) {
	c.mutex.Lock()
	// cash request interface to object
	request := data.(*checkout.CheckOutRequest)

	// initial domain object
	var transaction *domain.Transaction

	// initial response object
	var response *checkout.CheckoutResponse

	// check transaction exist
	checkTransaction, _ := c.transRepo.GetTransactionByOrderId(request.OrderId)

	if checkTransaction != nil {
		return nil, error2.OrderHasBeenExist
	}

	// get order
	order, err := c.orderRepo.GetOrderById(request.OrderId)
	if err != nil {
		return nil, error2.ErrNotFound
	}

	//check order exist
	if order != nil {
		var initialTransaction = domain.Transaction{}
		err := mapstructure.Decode(request, &initialTransaction)
		if err != nil {
			return nil, err
		}

		// create object transactionS
		initialTransaction.TransactionNumber = defaultFormatInvoice + strconv.Itoa(rand.Intn(100))
		initialTransaction.TransactionToken = uuid.New().String()
		initialTransaction.TransactionDate = time.Now().Add(time.Hour * 24)
		initialTransaction.Status = domain.Pending

		for _, orderItem := range order.OrderItems {
			// get product
			product, _ := c.productRepo.GetProductById(orderItem.ItemID)

			// check available
			if product.AvailableStock < orderItem.Quantity {
				return nil, error2.TotalStockLessThanQty
			}

			product.AvailableStock -= orderItem.Quantity
			product.StockOnHold += orderItem.Quantity

			go func() {
				// update stock product

				_, err := c.productRepo.UpdateProduct(product.ID, product)
				c.mutex.Unlock()
				if err != nil {
				}
			}()
		}

		// check checkout valid
		if request.Amount != order.TotalAmount {
			return nil, error2.CheckSumPaymentNotMatch
		}

		// create initial transaction
		transaction, err = c.transRepo.CreateTransaction(&initialTransaction)

		if err != nil {
			return nil, error2.ErrBadRequest
		}

		err2 := mapstructure.Decode(transaction, &response)
		if err2 != nil {
			return nil, err2
		}
		return response, nil

	} else {
		return nil, error2.ErrBadRequest
	}
}
