package routes

import (
	"alterra-agmc/day-4/config"
	"alterra-agmc/day-4/controllers"
	"alterra-agmc/day-4/handlers"
	"alterra-agmc/day-4/lib/database"
	"alterra-agmc/day-4/middlewares"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
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
	middlewares.CorsMiddleware(e)

	connection := config.Connect()
	userRepository := database.NewUserRepository(connection)
	userController := controllers.NewUserController(userRepository)

	// bookRepository := database.NewBookRepository(connection)
	// bookController := controllers.InitBook(bookRepository)

	v1 := e.Group("/api/v1")

	v1.POST("/users/", userController.CreateUserControllers)
	user := v1.Group("/users")
	// user.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningMethod: "HS512",
	// 	SigningKey:    []byte(os.Getenv("SECRET_KEY")),
	// 	TokenLookup:   "cookie:token",
	// }))
	user.GET("", userController.GetUsersControllers)
	user.GET("/:id", userController.GetUserByIdControllers)
	user.PUT("/:id", userController.UpdateUserControllers)
	user.DELETE("/:id", userController.DeleteUserControllers)

	// BOOK
	// v1.GET("/books/", bookController.GetBooksControllers)
	// v1.GET("/books/:id", bookController.GetBookByIdControllers)

	// book := v1.Group("/books")
	// book.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningMethod: "HS512",
	// 	SigningKey:    []byte(os.Getenv("SECRET_KEY")),
	// 	TokenLookup:   "cookie:token",
	// }))
	// book.POST("", bookController.CreateBookControllers)
	// book.PUT("/:id", bookController.UpdateBookControllers)
	// book.DELETE("/:id", bookController.DeleteBookControllers)
	return e
}
