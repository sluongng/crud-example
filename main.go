package main

import (
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"crud-example/handler"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.AddTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// Database connection
	db, err := sql.Open("mysql", "user:pass@tcp(localhost:3306)/user_db")
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()
	h := &handler.Handler{DB: db}

	// Root handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to User Service")
	})

	// User Service
	e.POST("/user",		h.Signup)
	e.GET("/user",			h.GetUserList)
	e.GET("/user/:id",		h.GetUser)
	e.PUT("/user",			h.UpdateUser)
	e.DELETE("/user/:id",	h.DeleteUser)

	e.Logger.Fatal(e.Start(":7001"))
}
