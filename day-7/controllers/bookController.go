package controllers

import (
	"alterra-agmc/day-7/handlers"
	"alterra-agmc/day-7/models"
	"alterra-agmc/day-7/repositories"
	"strconv"

	"github.com/labstack/echo"
)

type BookControllers struct {
	Lib repositories.BookRepository
}

type BookControllersInterface interface {
	GetBooksControllers(echo.Context) error
	GetBookByIdControllers(echo.Context) error
	CreateBookControllers(echo.Context) error
	UpdateBookControllers(echo.Context) error
	DeleteBookControllers(echo.Context) error
}

func InitBook(Lib repositories.BookRepository) BookControllersInterface {
	return &BookControllers{Lib}
}

func (controller BookControllers) GetBooksControllers(c echo.Context) error {
	data, err := controller.Lib.GetBook()

	if err != nil {
		handlers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}

	handlers.NewHandlerResponse("Successfully get books", data).Success(c)
	return nil
}

func (controller BookControllers) CreateBookControllers(c echo.Context) error {
	var book models.Book

	if err := c.Bind(&book); err != nil {
		handlers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}

	if err := handlers.DoValidation(&book); err != nil {
		handlers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return nil
	}

	controller.Lib.CreateBook(book)

	handlers.NewHandlerResponse("Successfully create book", nil).SuccessCreate(c)
	return nil
}

func (controller BookControllers) GetBookByIdControllers(c echo.Context) error {
	id := c.Param("id")
	idNumber, _ := strconv.Atoi(id)
	data, err := controller.Lib.GetBookByID(idNumber)

	if err != nil {
		handlers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}

	handlers.NewHandlerResponse("Successfully get book by id", data).Success(c)
	return nil
}

func (controller BookControllers) UpdateBookControllers(c echo.Context) error {
	id := c.Param("id")
	idNumber, _ := strconv.Atoi(id)
	var book models.Book

	if err := c.Bind(&book); err != nil {
		handlers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}

	if err := handlers.DoValidation(&book); err != nil {
		handlers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return nil
	}

	controller.Lib.UpdateBook(idNumber, book)

	handlers.NewHandlerResponse("Successfully update book by id", nil).Success(c)
	return nil
}

func (controller BookControllers) DeleteBookControllers(c echo.Context) error {
	id := c.Param("id")
	idNumber, _ := strconv.Atoi(id)

	_, err := controller.Lib.DeleteBook(idNumber)

	if err != nil {
		handlers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}

	handlers.NewHandlerResponse("Successfully delete book by id", nil).Success(c)
	return nil
}
