package handler

import (
	error2 "evermos_technical_test/error"
	"evermos_technical_test/pkg/dto/checkout"
	"evermos_technical_test/pkg/usecase"
	validator2 "evermos_technical_test/pkg/validator"
	"github.com/labstack/echo"
	"net/http"
)

type CheckOutHandler struct {
	CheckOutUseCase *usecase.CheckOutUseCase
}

func NewCheckOutHandler(e *echo.Echo, us *usecase.CheckOutUseCase) {
	handler := &CheckOutHandler{
		CheckOutUseCase: us,
	}
	// ENDPOINT
	e.POST("/transactions/checkouts", handler.CheckOutProcess)
}

func (p CheckOutHandler) CheckOutProcess(context echo.Context) (err error) {
	var request checkout.CheckOutRequest
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

	res, err := p.CheckOutUseCase.Process(&request)

	if err != nil {
		return context.JSON(error2.GetStatusCode(err), error2.ResponseError{Message: err.Error()})
	}

	return context.JSON(http.StatusCreated, res)
}
