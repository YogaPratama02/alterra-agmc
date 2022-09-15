package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

type HandlerResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewHandlerResponse(message string, data interface{}) *HandlerResponse {
	return &HandlerResponse{
		Message: message,
		Data:    data,
	}
}

func (response *HandlerResponse) Success(c echo.Context) {
	response.Status = true
	c.JSON(http.StatusOK, response)
}

func (response *HandlerResponse) SuccessCreate(c echo.Context) {
	response.Status = true
	c.JSON(http.StatusCreated, response)
}

func (response *HandlerResponse) Failed(c echo.Context) {
	response.Status = false
	c.JSON(http.StatusInternalServerError, response)
}
