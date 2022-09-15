package controllers

import (
	"alterra-agmc/day-2/handlers"
	"alterra-agmc/day-2/models"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

var (
	books = map[int]*models.Book{}
	seq   = 1
)

func GetBookAllControllers(c echo.Context) error {
	handlers.NewHandlerResponse("Successfully get all books", books).Success(c)
	return nil
}

func CreateBookControllers(c echo.Context) error {
	b := &models.Book{
		ID:        seq,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	c.Request().Header.Get("Content-type")

	if err := c.Bind(b); err != nil {
		return err
	}

	books[b.ID] = b
	seq++
	handlers.NewHandlerResponse("Successfully create books", b).SuccessCreate(c)
	return nil
	// return c.JSON(http.StatusCreated, b)
}

func GetBooksByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	handlers.NewHandlerResponse("Successfully get book by id", books[id]).Success(c)
	return nil
}

func UpdateBooksByID(c echo.Context) error {
	b := new(models.Book)
	if err := c.Bind(b); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	books[id].Title = b.Title
	books[id].Writter = b.Writter
	books[id].UpdatedAt = time.Now()
	return c.JSON(http.StatusOK, books[id])
}

func DeleteBooks(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(books, id)
	handlers.NewHandlerResponse("Successfully delete book by id", nil).Success(c)
	return nil
}
