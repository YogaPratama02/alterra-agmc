package routes

import (
	"alterra-agmc/day-2/config"
	"alterra-agmc/day-2/controllers"
	"alterra-agmc/day-2/lib/database"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()
	connection := config.Connect()

	userRepository := database.NewUserRepository(connection)
	userController := controllers.Init(userRepository)

	v1 := e.Group("/api/v1")

	user := v1.Group("/users")
	user.GET("", userController.GetUsersControllers)
	user.GET("/:id", userController.GetUserByIdControllers)
	user.POST("", userController.CreateUserControllers)
	user.PUT("/:id", userController.UpdateUserControllers)
	user.DELETE("/:id", userController.DeleteUserControllers)

	book := v1.Group("/books")
	book.GET("", controllers.GetBookAllControllers)
	book.GET("/:id", controllers.GetBooksByID)
	book.POST("", controllers.CreateBookControllers)
	book.PUT("/:id", controllers.UpdateBooksByID)
	book.DELETE("/:id", controllers.DeleteBooks)
	return e
}
