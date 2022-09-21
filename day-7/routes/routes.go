package routes

import (
	"alterra-agmc/day-7/config"
	"alterra-agmc/day-7/controllers"
	"alterra-agmc/day-7/handlers"
	"alterra-agmc/day-7/middlewares"
	"alterra-agmc/day-7/repositories"
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	// Using custom validation
	e.Validator = &handlers.CustomValidation{Validator: validator.New()}

	// Hide Echo Banner
	e.HideBanner = true

	// Middlewares
	middlewares.LoggerMiddleware(e)
	middlewares.RecoverMiddleware(e)

	connection := config.Connect()
	userRepository := repositories.NewUserRepository(connection)
	userController := controllers.Init(userRepository)

	bookRepository := repositories.NewBookRepository(connection)
	bookController := controllers.InitBook(bookRepository)

	v1 := e.Group("/api/v1")

	v1.POST("/users/", userController.CreateUserControllers)
	user := v1.Group("/users")
	user.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte(os.Getenv("SECRET_KEY")),
		TokenLookup:   "cookie:token",
	}))
	user.GET("", userController.GetUsersControllers)
	user.GET("/:id", userController.GetUserByIdControllers)
	user.PUT("/:id", userController.UpdateUserControllers)
	user.DELETE("/:id", userController.DeleteUserControllers)

	// BOOK
	v1.GET("/books/", bookController.GetBooksControllers)
	v1.GET("/books/:id", bookController.GetBookByIdControllers)

	book := v1.Group("/books")
	book.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte(os.Getenv("SECRET_KEY")),
		TokenLookup:   "cookie:token",
	}))
	book.POST("", bookController.CreateBookControllers)
	book.PUT("/:id", bookController.UpdateBookControllers)
	book.DELETE("/:id", bookController.DeleteBookControllers)
	return e
}
