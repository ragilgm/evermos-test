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
	"time"
)

type CheckOutUseCase struct {
	productRepo *repository.ProductRepository
	orderRepo   *repository.OrderRepository
	transRepo   *repository.TransactionRepository
}

func NewCheckOutUseCase(
	productRepo *repository.ProductRepository,
	orderRepo *repository.OrderRepository,
	transRepo *repository.TransactionRepository) *CheckOutUseCase {
	return &CheckOutUseCase{
		productRepo: productRepo,
		orderRepo:   orderRepo,
		transRepo:   transRepo,
	}
}

func (c CheckOutUseCase) Process(data interface{}) (interface{}, error) {

	min := 10
	max := 30

	request := data.(*checkout.CheckOutRequest)
	var transaction *domain.Transaction
	var response *checkout.CheckoutResponse

	// check transaction exist
	checkTransaction, _ := c.transRepo.GetTransactionByOrderId(request.OrderId)
	if checkTransaction != nil {
		if checkTransaction.Status == "SUCCESS" || checkTransaction.Status == "FAILED" {
			return nil, error2.ErrConflict
		}
	}

	// get order
	order, err := c.orderRepo.GetOrderById(request.OrderId)
	if err != nil {
		return nil, err
	}

	//check order exist
	if order != nil {
		var initialTransaction = domain.Transaction{}
		err := mapstructure.Decode(request, &initialTransaction)
		if err != nil {
			return nil, err
		}
		initialTransaction.TransactionNumber = "INV/" + strconv.Itoa(rand.Intn(max-min))
		initialTransaction.TransactionToken = uuid.New().String()
		initialTransaction.TransactionDate = time.Now().Add(time.Hour * 24)
		initialTransaction.Status = "PENDING"

		for _, orderItem := range order.OrderItems {
			// get product
			product, _ := c.productRepo.GetProductById(orderItem.ItemID)

			// check available
			if product.AvailableStock < orderItem.Quantity {
				return nil, error2.ErrBadRequest
			}

			product.AvailableStock -= orderItem.Quantity
			product.StockOnHold += orderItem.Quantity

			// update stock product
			go func() {
				_, err := c.productRepo.UpdateProduct(product.ID, product)
				if err != nil {

				}
			}()

		}

		// check checkout valid
		if request.Amount != order.TotalAmount {
			return nil, error2.ErrBadRequest
		}

		// create initial transaction
		transaction, _ = c.transRepo.CreateTransaction(&initialTransaction)

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
