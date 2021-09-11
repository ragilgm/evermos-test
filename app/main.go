package main

import (
	"evermos_technical_test/config"
	_repository "evermos_technical_test/pkg/repository"
	_usecase "evermos_technical_test/pkg/usecase"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"log"

	_handler "evermos_technical_test/pkg/handler"
	_middleware "evermos_technical_test/pkg/handler/middleware"
)

const port = ":8000"

func main() {

	e := echo.New()
	middle := _middleware.InitMiddleware()
	e.Use(middle.CORS)

	// initial repo
	productRepo := _repository.MewProductRepository(config.DbConfig{})
	orderRepo := _repository.NewOrderRepository(config.DbConfig{})
	transactionRepo := _repository.NewTransactionRepository(config.DbConfig{})

	// initial usecase
	au := _usecase.NewProductUseCase(productRepo)
	ou := _usecase.NewOrderUseCase(productRepo, orderRepo)
	pu := _usecase.NewPaymentUseCase(productRepo, orderRepo, transactionRepo)
	cu := _usecase.NewCheckOutUseCase(productRepo, orderRepo, transactionRepo)

	// initial handler
	_handler.NewProductHandler(e, au)
	_handler.NewOrderHandler(e, ou)
	_handler.NewPaymentHandler(e, pu)
	_handler.NewCheckOutHandler(e, cu)

	log.Fatal(e.Start(port))
}
