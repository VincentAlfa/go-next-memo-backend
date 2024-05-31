package main

import (
	"database/sql"
	"fmt"
	"go-next-memo/controller"
	"go-next-memo/database"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	dsn := database.DbSourceName()
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dsn.User, dsn.Password, dsn.Host, dsn.Port, dsn.Database))
	database.InitDB(db)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	

	e.Use(middleware.CORS())
	
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusAccepted, "string")
	})

	//users
	e.POST("/api/users/register", controller.RegisterUser)
	e.POST("/api/users/login", controller.LoginUser)




	e.Start(":4000")
}

