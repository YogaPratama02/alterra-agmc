package controllers

import (
	"alterra-agmc/day-3/handlers"
	"alterra-agmc/day-3/lib/database"
	"alterra-agmc/day-3/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type UserControllers struct {
	Lib database.UserRepository
}

type UserControllersInterface interface {
	GetUsersControllers(echo.Context) error
	GetUserByIdControllers(echo.Context) error
	CreateUserControllers(echo.Context) error
	UpdateUserControllers(echo.Context) error
	DeleteUserControllers(echo.Context) error
}

func Init(Lib database.UserRepository) UserControllersInterface {
	return &UserControllers{Lib}
}

func (controller UserControllers) GetUsersControllers(c echo.Context) error {
	data, err := controller.Lib.GetUser()

	if err != nil {
		handlers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}

	handlers.NewHandlerResponse("Successfully get users", data).Success(c)
	return nil
}

func (controller UserControllers) CreateUserControllers(c echo.Context) error {
	var user models.User
	var token models.Token

	if err := c.Bind(&user); err != nil {
		handlers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}

	if err := handlers.DoValidation(&user); err != nil {
		handlers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return nil
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 4)
	password := user.Password
	user.Password = string(hashPassword)

	userData, err := controller.Lib.CreateUser(user)
	if err != nil {
		handlers.NewHandlerValidationResponse(err.Error(), nil).BadRequest(c)
		return nil
	}

	errHash := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(password))
	if errHash != nil {
		log.Printf("Password Incorrect with err: %s\n", errHash)
		handlers.NewHandlerValidationResponse(errHash.Error(), nil).BadRequest(c)
		return nil
	}

	tokenString, err := handlers.GenerateJWT(userData)
	if err != nil {
		log.Printf("Error Generate JWT with err: %s\n", err)
		handlers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}
	token.Token = tokenString

	cookie := &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(30 * time.Hour),
		Path:     "/",
		SameSite: 2,
		HttpOnly: true,
	}
	c.SetCookie(cookie)
	handlers.NewHandlerResponse("Successfully create users", token).SuccessCreate(c)
	return nil
}

func (controller UserControllers) GetUserByIdControllers(c echo.Context) error {
	id := c.Param("id")
	idNumber, _ := strconv.Atoi(id)
	data, err := controller.Lib.GetUserByID(idNumber)

	if err != nil {
		handlers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}

	handlers.NewHandlerResponse("Successfully get user by id", data).Success(c)
	return nil
}

func (controller UserControllers) UpdateUserControllers(c echo.Context) error {
	id := c.Param("id")
	idNumber, _ := strconv.Atoi(id)
	var user models.User

	if err := c.Bind(&user); err != nil {
		handlers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}

	controller.Lib.UpdateUser(idNumber, user)

	handlers.NewHandlerResponse("Successfully update user by id", nil).Success(c)
	return nil
}

func (controller UserControllers) DeleteUserControllers(c echo.Context) error {
	id := c.Param("id")
	idNumber, _ := strconv.Atoi(id)

	_, err := controller.Lib.DeleteUser(idNumber)

	if err != nil {
		handlers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}

	handlers.NewHandlerResponse("Successfully delete user by id", nil).Success(c)
	return nil
}
