package validator 

import (
	validator "github.com/go-playground/validator/v10"
)

func NewValidator() *MyValidator {
	return &MyValidator{validator: validator.New()}
}

type MyValidator struct {
	validator *validator.Validate
}

func (cv *MyValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
