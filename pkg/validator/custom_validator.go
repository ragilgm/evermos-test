package validator

import (
	"github.com/go-playground/validator/v10"
)

// validator handler
func ValidateRequest(m interface{}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
