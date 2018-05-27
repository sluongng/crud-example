package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	sq "github.com/Masterminds/squirrel"

	"github.com/sluongng/crud-example/model"
)

// TODO: SignUp
// SignUp creates a user in database
func (h *Handler) SignUp(c echo.Context) (err error) {
	return c.String(http.StatusOK, "SignUp successful")
}

// GetUserList returns a list of Users
func (h *Handler) GetUserList(c echo.Context) (err error) {
	query, _, _ := sq.Select("*").
		From("user").
		Limit(100).
		ToSql()
	c.Logger().Info("Query to be executed is: %s", query)

	userList := new([]model.User)
	err = h.DB.Select(userList, query)
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
		c.Logger().Errorf("Invalid ID given: %d", userId)
		return err
	}

	query, _, _ := sq.Select("*").
		From("user").
		Where(sq.Eq{"id": userId}).
		Limit(1).
		ToSql()
	c.Logger().Info("Query to be executed is: %s", query)

	user := model.User{}
	err = h.DB.Get(&user, query, userId)
	if err != nil {
		c.Logger().Errorf("Could not find user with id: %d", userId)
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
