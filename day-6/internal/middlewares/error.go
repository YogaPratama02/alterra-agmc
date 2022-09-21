package middlewares

import (
	"net/http"

	res "alterra-agmc/day-6/pkg/utill/response"

	"github.com/labstack/echo"
)

func ErrorHandler(err error, c echo.Context) {
	var errCustom *res.Error

	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	switch report.Code {
	case http.StatusNotFound:
		errCustom = res.ErrorBuilder(&res.ErrorConstant.RouteNotFound, err)
	default:
		errCustom = res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	res.ErrorResponse(errCustom).Send(c)
}
