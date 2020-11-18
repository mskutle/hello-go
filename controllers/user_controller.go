package controllers

import (
	"encoding/json"
	"github.com/globalsign/mgo"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mskutle/hello-go/models"
	"github.com/mskutle/hello-go/services"
)

func GetAllUsers(c echo.Context) error {
	userService := services.NewUserService(c.Get("db").(*mgo.Database))
	err, users := userService.GetAll()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, users)
}

func AddUser(c echo.Context) error {
	var user models.User
	userService := services.NewUserService(c.Get("db").(*mgo.Database))

	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&user)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	if err := c.Validate(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	existingUser := userService.GetByUsername(user.Username)
	if existingUser != nil {
		return c.JSON(http.StatusBadRequest, models.NewErrorResponse("User with that username already exists"))
	}

	adduserErr, createdUser := userService.AddUser(user)

	if adduserErr != nil {
		return c.JSON(http.StatusInternalServerError, adduserErr.Error())
	}

	return c.JSON(http.StatusCreated, createdUser)
}

func GetUserByUsername(c echo.Context) error {
	userService := services.NewUserService(c.Get("db").(*mgo.Database))
	user := userService.GetByUsername(c.Param("username"))

	if user == nil {
		return c.JSON(http.StatusNotFound, nil)
	}

	return c.JSON(http.StatusOK, user)
}

func DeleteUserByUsername(c echo.Context) error {
	userService := services.NewUserService(c.Get("db").(*mgo.Database))
	user := userService.GetByUsername(c.Param("username"))

	if user == nil {
		return c.JSON(http.StatusNotFound, nil)
	}

	err := userService.Delete(user.Username)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.NewErrorResponse("An error occured"))
	}

	return c.JSON(http.StatusOK, "User deleted")
}