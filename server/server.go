package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mskutle/hello-go/controllers"
	customMiddleware "github.com/mskutle/hello-go/middleware"
)

type MyValidator struct {
	validator *validator.Validate
}

func (mv *MyValidator) Validate(i interface{}) error {
	return mv.validator.Struct(i)
}

func StartServer(port string) *echo.Echo {
	server := echo.New()
	server.Validator = &MyValidator{validator: validator.New()}
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	// server.Use(customMiddleware.MongoDb())

	server.GET("/", controllers.Root)

	server.GET("/users", controllers.GetAllUsers)
	server.GET("/users/:username", controllers.GetUserByUsername)
	server.DELETE("/users/:username", controllers.DeleteUserByUsername)
	server.POST("/users", controllers.AddUser)

	server.POST("/login", controllers.Login)

	server.Logger.Fatal(server.Start(port))

	return server
}
