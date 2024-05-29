package controller

import (
	"go-next-memo/models"
	"net/http"
	"go-next-memo/database"
	"github.com/labstack/echo/v4"
)


func SelectALLUser(c echo.Context) error {
	db := database.GetDB()
	row, err := db.Query("SELECT email, password FROM user")
	if err != nil {
		panic(err.Error())
	}
	defer row.Close()

	var users []model.User

	for row.Next() {
		var user model.User
		err := row.Scan(&user.Email, &user.Password)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}
	return c.JSON(http.StatusAccepted, users)
}