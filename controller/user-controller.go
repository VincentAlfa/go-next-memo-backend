package controller

import (
	"go-next-memo/database"
	"go-next-memo/models"
	"go-next-memo/utils"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

var sessionId = "sessions"
var store = sessions.NewCookieStore([]byte(sessionId))

func AuthMiddleWare (next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, _ := store.Get(c.Request(), sessionId)

		if session.Values["email"] == nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Unauthorized"})
		}
		return next(c)
	}
}


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


	session, _ := store.Get(c.Request(), sessionId)
	session.Values["email"] = user.Email
	session.Save(c.Request(), c.Response())  

	row.Scan(&user.Email, &user.Password)
	return c.JSON(http.StatusAccepted, echo.Map{"message" : "login Success" })
}

// func ForgotPasswordUser(c echo.Context) error {
// 	db := database.GetDB()
// 	user := model.User{}
// 	err := c.Bind(&user)
// 	if err != nil {
// 		panic(err)
// 	}
// 	query := 
// 	row := db.Exec()

// }