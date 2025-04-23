package util

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(s interface{}) map[string]string {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	errors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		field := strings.ToLower(err.Field())
		errors[field] = getErrorMsg(err.Tag())
	}
	return errors
}

func getErrorMsg(tag string) string {
	switch tag {
	case "required":
		return "wajib diisi"
	case "email":
		return "format email tidak valid"
	case "min":
		return "terlalu pendek"
	default:
		return "tidak valid"
	}
}
