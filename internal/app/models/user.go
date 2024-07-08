package models

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password,omitempty" validate:"min=8"`
	Avatar    string    `json:"avatar" validate:"url"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserPublic represents the public user data
type UserPublic struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
	Bio       string `json:"bio"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (u *User) Validate() error {
	validate := validator.New()

	if err := validate.StructExcept(u, "Password"); err != nil {
		return err
	}

	return nil
}

func (u *User) ValidateCreate() error {
	validate := validator.New()

	if err := validate.Struct(u); err != nil {
		return err
	}

	if len(u.Password) < 8 {
		return errors.New("password is required and must be at least 8 characters long")
	}

	return nil
}
