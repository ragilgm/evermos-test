package handler

import (
	_ "context"
	error2 "evermos_technical_test/error"
	product2 "evermos_technical_test/pkg/dto/product"
	"evermos_technical_test/pkg/usecase"
	validator2 "evermos_technical_test/pkg/validator"
	_ "fmt"
	"github.com/labstack/echo"
	_ "github.com/mitchellh/mapstructure"
	"net/http"
	"strconv"
)

// ProductHandler  represent the httpHandler for product
type ProductHandler struct {
	pUCase *usecase.ProductUseCase
}

// NewProductHandler will initialize the products/ resources endpoint
func NewProductHandler(e *echo.Echo, us *usecase.ProductUseCase) {
	handler := &ProductHandler{
		pUCase: us,
	}
	// ENDPOINT
	e.GET("/products", handler.GetProducts)
	e.GET("/products/:id", handler.GetProductById)
	e.PUT("/products/:id", handler.UpdateProduct)
	e.POST("/products", handler.CreateProduct)
	e.DELETE("/products/:id", handler.DeleteProduct)
}

// get all products
func (p ProductHandler) GetProducts(context echo.Context) (err error) {

	// useCase flow
	listProducts, err := p.pUCase.GetProducts()

	if err != nil {
		return context.JSON(error2.GetStatusCode(err), error2.ResponseError{Message: err.Error()})
	}
	return context.JSON(http.StatusOK, listProducts)
}

// get product by id
func (p ProductHandler) GetProductById(context echo.Context) (err error) {
	var id, _ = strconv.Atoi(context.Param("id"))

	// useCase flow
	listProducts, err := p.pUCase.GetProductById(uint(id))

	if err != nil {
		return context.JSON(error2.GetStatusCode(err), error2.ResponseError{Message: err.Error()})
	}
	return context.JSON(http.StatusOK, listProducts)
}

// update product
func (p ProductHandler) UpdateProduct(context echo.Context) (err error) {
	var request product2.ProductRequest

	// collect request
	err = context.Bind(&request)

	// check request null
	if err != nil {
		return context.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	// collect id
	var id, _ = strconv.Atoi(context.Param("id"))

	// validate request
	var ok bool
	if ok, err = validator2.ValidateRequest(&request); !ok {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	// useCase flow
	res, err := p.pUCase.UpdateProduct(uint(id), &request)

	if err != nil {
		return context.JSON(error2.GetStatusCode(err), error2.ResponseError{Message: err.Error()})
	}

	return context.JSON(http.StatusCreated, res)
}

// delete product
func (p ProductHandler) DeleteProduct(context echo.Context) error {

	// collect id
	var id, _ = strconv.Atoi(context.Param("id"))

	// useCase flow
	err := p.pUCase.DeleteProduct(int64(id))

	if err != nil {
		return context.JSON(error2.GetStatusCode(err), error2.ResponseError{Message: err.Error()})
	}
	return context.JSON(http.StatusNoContent, nil)
}

func (p ProductHandler) CreateProduct(context echo.Context) (err error) {
	var request product2.ProductRequest

	// collect request
	err = context.Bind(&request)

	// check request null
	if err != nil {
		return context.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	// check validate request
	var ok bool
	if ok, err = validator2.ValidateRequest(&request); !ok {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	// call useCase
	res, err := p.pUCase.CreateProduct(&request)

	if err != nil {
		return context.JSON(error2.GetStatusCode(err), error2.ResponseError{Message: err.Error()})
	}

	return context.JSON(http.StatusCreated, res)
}
