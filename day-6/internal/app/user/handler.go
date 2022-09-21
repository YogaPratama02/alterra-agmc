package user

import (
	"alterra-agmc/day-6/internal/factory"

	"alterra-agmc/day-6/internal/dto"
	res "alterra-agmc/day-6/pkg/utill/response"

	"github.com/labstack/echo"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}

func (h *handler) Create(c echo.Context) error {
	payload := new(dto.UserPayload)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		response := res.ErrorBuilder(&res.ErrorConstant.Validation, err)
		return response.Send(c)
	}

	err := h.service.CreateUser(c.Request().Context(), payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(nil).Send(c)
}
