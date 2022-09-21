package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

type HandlerValidationResponse struct {
	Status  bool        `json:"status"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

func NewHandlerValidationResponse(message, data interface{}) *HandlerValidationResponse {
	return &HandlerValidationResponse{
		Message: message,
		Data:    data,
	}
}

func (response *HandlerValidationResponse) Failed(c echo.Context) {
	response.Status = false
	c.JSON(http.StatusInternalServerError, response)
}

func (response *HandlerValidationResponse) BadRequest(c echo.Context) {
	response.Status = false
	c.JSON(http.StatusBadRequest, response)
}
