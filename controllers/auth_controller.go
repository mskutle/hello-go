package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mskutle/hello-go/services"
)

type login struct {
	username string `json: "username" validate: "nonzero"`
	password string `json: "password" validate: "nonzero"`
}

func Login(c echo.Context) error {
	var input login
	userService := c.Get("userService").(services.UserService)
	defer c.Request().Body.Close()
	fmt.Print("input", c.Request())

	if err := json.NewDecoder(c.Request().Body).Decode(&input); err != nil {
		return err
	}

	if err := c.Validate(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}
	user := userService.GetByUsername(input.username)
	fmt.Println(user)
	if user == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	if input.password == user.Password {
		return c.JSON(http.StatusOK, user)
	}
	return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized.")
}
