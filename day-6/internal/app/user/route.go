package user

import "github.com/labstack/echo"

func (h *handler) Route(g *echo.Group) {
	g.POST("/api/v1/users", h.Create)
}
