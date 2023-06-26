package validator

import "github.com/go-playground/validator/v10"

var validate = validator.New()

func ValidateStruct[T any](payload T) error {
	err := validate.Struct(payload)
	return err
}
