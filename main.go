package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/go-sql-driver/mysql"

	"crud-example/handler"
	"database/sql"
)

func rootHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to User Service")
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.AddTrailingSlash())
	e.Use(middleware.Recover())

	// Database connection
	// TODO: move DataSource name to flag()
	db, err := sql.Open("mysql", "root:root@MyUserDb")
	if err != nil {
		e.Logger.Fatal(err)
	}

	// TODO: ensure migration is ran

	h := &handler.Handler{DB: db}

	e.GET("/", 		rootHandler)

	e.GET("/user",			h.GetUserList)
	e.GET("/user/:id",		h.GetUser)
	e.POST("/user",		h.Signup)
	e.PUT("/user",			h.UpdateUser)

	e.Logger.Fatal(e.Start(":7001"))
}
