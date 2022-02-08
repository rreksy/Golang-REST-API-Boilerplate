package helper

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		var errorsMessage string
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				errorsMessage = fmt.Sprintf("%s is required",
					err.Field())
			case "email":
				errorsMessage = fmt.Sprintf("%s is not valid email",
					err.Field())
			case "gte":
				errorsMessage = fmt.Sprintf("%s value must be greater than %s",
					err.Field(), err.Param())
			case "lte":
				errorsMessage = fmt.Sprintf("%s value must be lower than %s",
					err.Field(), err.Param())
			}
			errors = append(errors, errorsMessage)
			// break
		}
	}

	return errors
}

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")

	if err != nil {
		return "", err
	}

	return string(val), nil
}
