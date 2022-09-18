package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type CustomValidation struct {
	Validator *validator.Validate
}

func (cv *CustomValidation) Validate(i interface{}) error {

	// Check Ada Validator Error Atau Enggak
	if err := cv.Validator.Struct(i); err != nil {
		errorMessage := ErrorMessage(err)
		return echo.NewHTTPError(http.StatusBadRequest, errorMessage)
	}

	return nil
}

func ErrorMessage(check error) map[string]string {
	var message map[string]string

	message = map[string]string{}

	for _, e := range check.(validator.ValidationErrors) {
		switch e.Tag() {
		case "required":
			message[strings.ToLower(e.Field())] = fmt.Sprintf("Field %s can not empty!", e.Field())
		case "email":
			message[strings.ToLower(e.Field())] = fmt.Sprintf("Input must be email!")
		case "datetime":
			message[strings.ToLower(e.Field())] = fmt.Sprintf("Field %s must be date & time", e.Field())
		}
	}

	return message
}

func DoValidation(i interface{}) map[string]string {
	var message map[string]string

	message = map[string]string{}

	val := CustomValidation{validator.New()}

	if err := val.Validator.Struct(i); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			switch e.Tag() {
			case "required":
				message[strings.ToLower(e.Field())] = fmt.Sprintf("Field %s can not empty!", e.Field())
			case "email":
				message[strings.ToLower(e.Field())] = fmt.Sprintf("Input must be email!")
			case "datetime":
				message[strings.ToLower(e.Field())] = fmt.Sprintf("Field %s must be date & time", e.Field())
			}
		}

		return message
	}

	return nil
}
