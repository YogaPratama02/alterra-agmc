package controllers

import (
	"alterra-agmc/day-2/handlers"
	"alterra-agmc/day-2/lib/database"
	"alterra-agmc/day-2/models"
	"strconv"

	"github.com/labstack/echo"
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

	if err := c.Bind(&user); err != nil {
		handlers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}

	controller.Lib.CreateUser(user)

	handlers.NewHandlerResponse("Successfully create users", nil).SuccessCreate(c)
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
