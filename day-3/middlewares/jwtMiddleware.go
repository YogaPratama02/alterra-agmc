package middlewares

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetJwtMiddlewares(g *echo.Group) {
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte(os.Getenv("JWT_SECRET")),
		TokenLookup:   "cookie:token",
	}))
	middleware.ErrJWTMissing.Code = 401
	middleware.ErrJWTMissing.Message = "You're unauthorized"
}
