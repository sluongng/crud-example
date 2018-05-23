package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func rootHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Wellcome to A sluong's experiment")
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", rootHandler)
	e.GET("/hello", hello)

	e.Logger.Fatal(e.Start(":7001"))
}
