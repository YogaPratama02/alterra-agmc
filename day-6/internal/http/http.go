package http

import (
	"alterra-agmc/day-6/internal/app/user"
	"alterra-agmc/day-6/internal/factory"

	"github.com/labstack/echo"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	user.NewHandler(f).Route(e.Group("/products"))
}
