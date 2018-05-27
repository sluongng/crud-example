package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/sluongng/crud-example/handler"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.Logger())
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// Database connection
	db, err := sql.Open("mysql", "user:pass@tcp(database:3306)/user_db")
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
	u := e.Group("/user")

	u.POST("", h.Signup)
	u.GET("", h.GetUserList)
	u.GET("/:id", h.GetUser)
	u.PUT("/:id", h.UpdateUser)
	u.DELETE("/:id", h.DeleteUser)

	e.Logger.Fatal(e.Start(":7001"))
}
