package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"crud-example/handler"
)

func rootHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to A sluong's experiment")
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	h := &handler.Handler{}

	e.GET("/", 		rootHandler)

	e.GET("/user",			h.GetUserList)
	e.GET("/user/:id",		h.GetUser)
	e.POST("/user",		h.Signup)
	e.PUT("/user",			h.UpdateUser)

	e.Logger.Fatal(e.Start(":7001"))
}
