package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidationStruct(input interface{}) []string {
	var validationError []string

	err := validator.New().Struct(input)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			message := fmt.Sprintf("%s is %s", err.Field(), err.ActualTag())
			validationError = append(validationError, message)
		}
	}

	return validationError
}
