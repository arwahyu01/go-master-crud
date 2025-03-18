package helpers

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

var validate = validator.New()

type ValidationError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Error string `json:"error"`
}

func ValidateStruct(s interface{}) []ValidationError {
	var errors []ValidationError
	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, ValidationError{
				Field: err.Field(),
				Tag:   err.Tag(),
				Error: strings.ToLower(err.Field()) + " is " + err.Tag(),
			})
		}
	}
	return errors
}
