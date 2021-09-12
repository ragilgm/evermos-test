package main

import (
	"evermos_technical_test/config"
	_repository "evermos_technical_test/pkg/repository"
	_usecase "evermos_technical_test/pkg/usecase"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"log"
	"sync"

	_handler "evermos_technical_test/pkg/handler"
	_middleware "evermos_technical_test/pkg/handler/middleware"
)

const port = ":8000"

func main() {

	e := echo.New()
	middle := _middleware.InitMiddleware()
	e.Use(middle.CORS)

	var mutex sync.Mutex

	// initial repo
	productRepo := _repository.MewProductRepository(config.DbConfig{})
	orderRepo := _repository.NewOrderRepository(config.DbConfig{})
	transactionRepo := _repository.NewTransactionRepository(config.DbConfig{})

	// initial usecase
	au := _usecase.NewProductUseCase(mutex, productRepo)
	ou := _usecase.NewOrderUseCase(mutex, productRepo, orderRepo)
	pu := _usecase.NewPaymentUseCase(mutex, productRepo, orderRepo, transactionRepo)
	cu := _usecase.NewCheckOutUseCase(mutex, productRepo, orderRepo, transactionRepo)

	// initial handler
	_handler.NewProductHandler(e, au)
	_handler.NewOrderHandler(e, ou)
	_handler.NewPaymentHandler(e, pu)
	_handler.NewCheckOutHandler(e, cu)

	log.Fatal(e.Start(port))
}
