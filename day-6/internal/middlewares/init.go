package middlewares

import (
	"alterra-agmc/day-6/pkg/utill/validator"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	e.HTTPErrorHandler = ErrorHandler
	e.Validator = &validator.CustomValidator{Validator: validator.NewValidator()}

}
