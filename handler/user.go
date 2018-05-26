package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"fmt"
	"crud-example/model"
)

func (h *Handler) Signup(c echo.Context) (err error) {
	return c.String(http.StatusOK, "Signup successful")
}

// GetUserList returns a list of Users
func (h *Handler) GetUserList(c echo.Context) (err error) {
	return c.String(http.StatusOK, "Get some User")
}

// GetUser returns a User
func (h *Handler) GetUser(c echo.Context) (err error) {
	userId := c.Param("id")

	var user model.User
	err = h.DB.QueryRow("SELECT * FROM user WHERE id = ?", userId).
		Scan(&user.ID, &user.Name, &user.Email, &user.Website)

	if err != nil {
		c.Logger().Errorf("Could not find user with id: %s", userId)
		return err
	}

	return c.String(http.StatusOK, fmt.Sprintf("Username is %s and email is %s", user.Name, user.Email))
}

// UpdateUser returns result of User update
func (h *Handler) UpdateUser(c echo.Context) (err error) {
	return c.String(http.StatusOK, "Updated User successful")
}

// DeleteUser delete a user by id
func (h *Handler) DeleteUser(c echo.Context) (err error) {
	return c.String(http.StatusOK, "Deleted User successful")
}
