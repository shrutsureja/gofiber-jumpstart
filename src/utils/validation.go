package utils

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ValidateStruct validates a given struct and returns structured error messages
func ValidateStruct(data interface{}, validationMessage map[string]string) (map[string]string, error) {
	err := validate.Struct(data)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make(map[string]string)
		for _, err := range validationErrors {
			fieldTag := err.Field() + "." + err.Tag()
			message, ok := validationMessage[fieldTag]
			if ok == false {
				message = validationMessage[err.Field()]
			}
			if message == "" {
				message = "Invalid InputValue"
			}
			errorMessages[err.Field()] = message
		}
		return errorMessages, err
	}
	return nil, nil
}
