package main

import (
	_ "github.com/go-sql-driver/mysql"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"

	"github.com/sluongng/crud-example/handler"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetLevel(log.DEBUG)

	e.Use(middleware.Logger())
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// Database connection
	dataSourceName := "user:pass@tcp(database:3306)/user_db"
	db, err := sqlx.Open("mysql", dataSourceName)
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

	u.POST("", h.SignUp)
	u.GET("", h.GetUserList)
	u.GET("/:id", h.GetUser)
	u.PUT("/:id", h.UpdateUser)
	u.DELETE("/:id", h.DeleteUser)

	e.Logger.Fatal(e.Start(":7001"))
}
