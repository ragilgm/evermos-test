package handler

import (
	_ "context"
	error2 "evermos_technical_test/error"
	"evermos_technical_test/pkg/dto/order"
	"evermos_technical_test/pkg/usecase"
	validator2 "evermos_technical_test/pkg/validator"
	_ "fmt"
	"github.com/labstack/echo"
	_ "github.com/mitchellh/mapstructure"
	"net/http"
)

type OrderHandler struct {
	oUcase *usecase.OrderUseCase
}

func NewOrderHandler(e *echo.Echo, us *usecase.OrderUseCase) {
	handler := &OrderHandler{
		oUcase: us,
	}
	// ENDPOINT
	e.POST("/orders", handler.CreateOrder)

}

func (p OrderHandler) CreateOrder(context echo.Context) (err error) {
	var request order.OrderRequest

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
	res, err := p.oUcase.CreateOrder(&request)

	if err != nil {
		return context.JSON(error2.GetStatusCode(err), error2.ResponseError{Message: err.Error()})
	}

	return context.JSON(http.StatusCreated, res)
}
