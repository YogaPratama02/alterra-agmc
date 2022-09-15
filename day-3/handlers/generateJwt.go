package handlers

import (
	"alterra-agmc/day-3/models"
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

type JWTClaim struct {
	jwt.StandardClaims
}

var SecretKey = []byte(os.Getenv("SECRET_KEY"))

func GenerateJWT(user *models.User) (string, error) {
	expirationTime := time.Now().Add(30 * time.Hour)
	claims := &JWTClaim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    strconv.Itoa(user.Id),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Printf("Error signed string jwt with err: %s\n", err)
		return tokenString, err
	}
	return tokenString, err
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}

func ParseJWTClaims(c echo.Context) (*int, error) {
	var userId interface{}

	token, err := c.Cookie("token")
	if err != nil {
		log.Printf("Error get cookie with err: %s", err)
		return nil, err
	}
	claims := jwt.MapClaims{}
	res, err := jwt.ParseWithClaims(token.Value, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		log.Printf("Error parse jwt token with err: %s", err)
		return nil, err
	}
	resMap, ok := res.Claims.(jwt.MapClaims)
	if ok && res.Valid {
		userId = resMap["iss"]
		convertUserId, _ := strconv.Atoi(userId.(string))
		return &convertUserId, err
	} else {
		log.Printf("Invalid token jwt")
		return nil, err
	}
}
