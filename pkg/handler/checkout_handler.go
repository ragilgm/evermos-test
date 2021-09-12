package handler

import (
	error2 "evermos_technical_test/error"
	"evermos_technical_test/pkg/dto/checkout"
	"evermos_technical_test/pkg/usecase"
	validator2 "evermos_technical_test/pkg/validator"
	"github.com/labstack/echo"
	"net/http"
	"sync"
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
	var response *checkout.CheckoutResponse
	var errStatus = http.StatusCreated
	var waitGroup sync.WaitGroup
	var mutex sync.Mutex
	waitGroup.Add(1)

	go func() {
		// collect request
		err = context.Bind(&request)

		// check validate request
		var ok bool

		mutex.Lock()
		ok, err1 := validator2.ValidateRequest(&request)
		if !ok {
			errStatus = http.StatusBadRequest
			err = err1

			// set wait group done
			waitGroup.Done()
			return
		}
		mutex.Unlock()

		mutex.Lock()
		res, err2 := p.CheckOutUseCase.Process(&request)
		if err2 != nil {
			err = err2
			errStatus = error2.GetStatusCode(err)

			// set wait group done
			waitGroup.Done()
			return
		}
		if res != nil {
			response = res.(*checkout.CheckoutResponse)
		}
		mutex.Unlock()

		// set wait group done
		waitGroup.Done()

	}()

	// wait wait grop done
	waitGroup.Wait()

	if err != nil {
		return context.JSON(errStatus, error2.ResponseError{Message: err.Error()})
	}

	return context.JSON(http.StatusCreated, response)
}
