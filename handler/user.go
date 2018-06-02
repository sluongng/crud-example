package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	sq "github.com/Masterminds/squirrel"

	"github.com/sluongng/crud-example/model"
)

// TODO: SignUp

// SignUp godoc
// @Summary create user
// @Description create a new user to DB
// @Tags User
// @Produce plain
// @Success 200 {string} string
// @Router /user [post]
func (h *Handler) SignUp(c echo.Context) (err error) {
	return c.String(http.StatusOK, "SignUp successful")
}

// GetUserList godoc
// @Summary get list of users
// @Description get list of current users from Database
// @Tags User
// @Produce json
// @Success 200 {array} model.User
// @Router /user [get]
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

// GetUser godoc
// @Summary get 1 user detail
// @Description Get a user by id
// @Tags User
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} model.User
// @Router /user/{id} [get]
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
// UpdateUser godoc
// @Summary Update a user's detail
// @Description update a user by id
// @Tags User
// @Produce plain
// @Param id path int true "User ID"
// @Success 200 {string} string
// @Router /user/{id} [put]
func (h *Handler) UpdateUser(c echo.Context) (err error) {
	return c.String(http.StatusOK, "Updated User successful")
}

// TODO: DeleteUser
// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user by id
// @Tags User
// @Produce plain
// @Param id path int true "User ID"
// @Success 200 {string} string
// @Router /user/{id} [delete]
func (h *Handler) DeleteUser(c echo.Context) (err error) {
	return c.String(http.StatusOK, "Deleted User successful")
}
