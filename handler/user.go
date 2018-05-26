package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

// Signup godoc
// @Summary Creating new user
// @Description Create new user
// @Tags user
// @Accept json
// @Produce json
// @Param user body model.NewUser true "user info"
// @Success 200 {object} model.User
// @Router /user [post]
func (h *Handler) Signup(c echo.Context) (err error) {
	return c.String(http.StatusOK, "Signup successful")
}

// GetUserList returns a list of Users
func (h *Handler) GetUserList(c echo.Context) (err error) {
	return c.String(http.StatusOK, "Get some User")
}

// GetUser returns a User
func (h *Handler) GetUser(c echo.Context) (err error) {
	return c.String(http.StatusOK, "Get one User")
}

// UpdateUser returns result of User update
func (h *Handler) UpdateUser(c echo.Context) (err error) {
	return c.String(http.StatusOK, "Updated User successful")
}
