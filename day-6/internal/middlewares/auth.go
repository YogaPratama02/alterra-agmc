package middlewares

import (
	res "alterra-agmc/day-6/pkg/utill/response"
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	var (
		jwtKey = os.Getenv("JWT_KEY")
	)

	return func(c echo.Context) error {
		authToken := c.Request().Header.Get("Authorization")
		if authToken == "" {
			return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, nil).Send(c)
		}

		splitToken := strings.Split(authToken, "Bearer ")
		token, err := jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method :%v", token.Header["alg"])
			}

			return []byte(jwtKey), nil
		})

		if !token.Valid || err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err).Send(c)
		}

		// todo add to context

		// var id uint
		// ID := token.Claims.(jwt.MapClaims)["id"]
		// if ID != nil {
		// 	id = ID.(uint)
		// } else {
		// 	if err != nil {
		// 		return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err).Send(c)
		// 	}
		// }

		return next(c)
	}
}
