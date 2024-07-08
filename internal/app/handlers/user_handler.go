package handlers

import (
	"net/http"
	"strconv"

	"github.com/jonathanqueiroz/tickets/internal/app/models"
	"github.com/jonathanqueiroz/tickets/internal/app/repositories"
	"github.com/labstack/echo"
)

type UserHandler struct {
	Repo repositories.UserRepositoryInterface
}

func NewUserHandler(repo repositories.UserRepositoryInterface) *UserHandler {
	return &UserHandler{
		Repo: repo,
	}
}

// CreateUser creates a new user
func (h *UserHandler) CreateUser(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	if err := user.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	createdUser, err := h.Repo.Create(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, createdUser)
}

// GetUsers gets a user by ID
func (h *UserHandler) GetUsers(c echo.Context) error {
	users, err := h.Repo.GetUsers()
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, users)
}

// GetUser gets a user by ID
func (h *UserHandler) GetUser(c echo.Context) error {
	user, err := h.getUserByID(c)
	if err != nil {
		httpError := err.(*echo.HTTPError)
		return c.JSON(httpError.Code, map[string]string{"message": httpError.Message.(string)})
	}

	return c.JSON(http.StatusOK, user)
}

// UpdateUser updates a user by ID
func (h *UserHandler) UpdateUser(c echo.Context) error {
	userSaved, err := h.getUserByID(c)
	if err != nil {
		httpError := err.(*echo.HTTPError)
		return c.JSON(httpError.Code, map[string]string{"message": httpError.Message.(string)})
	}

	var user models.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	if err := user.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	err = h.Repo.Update(userSaved.ID, &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.NoContent(http.StatusOK)
}

// DeleteUser deletes a user by ID
func (h *UserHandler) DeleteUser(c echo.Context) error {
	user, err := h.getUserByID(c)
	if err != nil {
		httpError := err.(*echo.HTTPError)
		return c.JSON(httpError.Code, map[string]string{"message": httpError.Message.(string)})
	}

	err = h.Repo.Delete(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.NoContent(http.StatusOK)
}

func (h *UserHandler) getUserByID(c echo.Context) (*models.UserPublic, error) {
	userIdStr := c.Param("id")

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	user, err := h.Repo.GetByID(userId)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return user, nil
}
