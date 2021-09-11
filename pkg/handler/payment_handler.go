package handler

import (
	error2 "evermos_technical_test/error"
	"evermos_technical_test/pkg/dto/payment"
	"evermos_technical_test/pkg/usecase"
	validator2 "evermos_technical_test/pkg/validator"
	"github.com/labstack/echo"
	"net/http"
)

type PaymentHandler struct {
	paymentUseCase *usecase.PaymentUseCase
}

func NewPaymentHandler(e *echo.Echo, us *usecase.PaymentUseCase) {
	handler := &PaymentHandler{
		paymentUseCase: us,
	}
	// ENDPOINT
	e.PATCH("/transactions/payments", handler.PaymentProcess)
}

func (p PaymentHandler) PaymentProcess(context echo.Context) (err error) {
	var request payment.PaymentRequest

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
	res, err := p.paymentUseCase.Process(&request)

	if err != nil {
		return context.JSON(error2.GetStatusCode(err), error2.ResponseError{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, res)
}
