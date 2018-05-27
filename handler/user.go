package handler

import (
	"github.com/labstack/echo"
	"github.com/sluongng/crud-example/model"
	"net/http"
	"strconv"
)

// TODO: Signup
// Signup creates a user in database
func (h *Handler) Signup(c echo.Context) (err error) {
	return c.String(http.StatusOK, "Signup successful")
}

// GetUserList returns a list of Users
func (h *Handler) GetUserList(c echo.Context) (err error) {

	userList := new([]model.User)
	err = h.DB.Select(userList, "SELECT * FROM user LIMIT 100")
	if err != nil {
		c.Logger().Errorf("Error getting user list")
		return err
	}

	return c.JSON(http.StatusOK, userList)
}

// GetUser returns a User
func (h *Handler) GetUser(c echo.Context) (err error) {

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Errorf("Invalid ID given: %s", userId)
		return err
	}

	user := new(model.User)
	err = h.DB.Get(&user, "SELECT * FROM user WHERE id = ?", userId)
	if err != nil {
		c.Logger().Errorf("Could not find user with id: %s", userId)
		return err
	}

	return c.JSON(http.StatusOK, user)
}

// TODO: UpdateUser
// UpdateUser returns result of User update
func (h *Handler) UpdateUser(c echo.Context) (err error) {
	return c.String(http.StatusOK, "Updated User successful")
}

// TODO: DeleteUser
// DeleteUser delete a user by id
func (h *Handler) DeleteUser(c echo.Context) (err error) {
	return c.String(http.StatusOK, "Deleted User successful")
}
