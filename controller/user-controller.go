package controller

import (
	"go-next-memo/database"
	"go-next-memo/models"
	"go-next-memo/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterUser(c echo.Context) error {
	db := database.GetDB()
	user := model.User{}
	err := c.Bind(&user)
	if err != nil {
		panic(err)
	}
	
	userData, _ := utils.GetUserByEmail(user.Email)

	if user.Email == userData.Email {
		return c.JSON(http.StatusBadRequest, echo.Map{"message" : "email already registered invalid"})
	} 

	if user.Email == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "email is empty"})
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		panic(err)
	}

	query := "INSERT INTO user (email, password) values (?, ?)"
	_, err = db.Exec(query, user.Email, hashedPassword)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request"})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "register success"})
}

func LoginUser (c echo.Context) error {
	db := database.GetDB()
	user := model.User{}
	err := c.Bind(&user)
	if err != nil {
		panic(err)
	}

	userData, _ := utils.GetUserByEmail(user.Email)

	if user.Email != userData.Email {
		return c.JSON(http.StatusBadRequest, echo.Map{"message" : "email invalid"})
	} 

	res := utils.CheckHashedPassword(user.Password, userData.Password)
	if !res {
		return c.JSON(http.StatusBadRequest, echo.Map{"message" : "password invalid"})
	}
	
	query := "SELECT email, password FROM user WHERE email = ? "
	row := db.QueryRow(query, user.Email)
	if row.Err() != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message" : row.Err()})
	}

	row.Scan(&user.Email, &user.Password)
	return c.JSON(http.StatusAccepted, echo.Map{"message" : "login Success" })
}