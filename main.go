package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/swaggo/echo-swagger"
	_ "github.com/sluongng/crud-example/docs"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/sluongng/crud-example/handler"
)

func rootHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to User Service")
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server User-service.

// @contact.name Son Luong Ngoc
// @contact.url https://sluongng.gitlab.io/
// @contact.email sluongng@gmail.com
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

	// Root handler
	e.GET("/", 		rootHandler)

	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// User Service
	e.GET("/user",			h.GetUserList)
	e.GET("/user/:id",		h.GetUser)
	e.POST("/user",		h.Signup)
	e.PUT("/user",			h.UpdateUser)

	e.Logger.Fatal(e.Start(":7001"))
}
