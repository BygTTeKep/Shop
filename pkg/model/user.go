package model

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type User struct {
	Id            int    `json:"id"`
	UserName      string `json:"username"`
	FirstName     string `json:"firstname"`
	LastName      string `json:"lastname"`
	Surname       string `json:"surname"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	password_hash string `json:"-"`
}

type UpdateUser struct {
	UserName      *string `json:"username"`
	FirstName     *string `json:"firstname"`
	LastName      *string `json:"lastname"`
	Surname       *string `json:"surname"`
	Email         *string `json:"email"`
	Password      *string `json:"password"`
	password_hash *string `json:"-"`
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.UserName, validation.Required),
		validation.Field(&u.FirstName, validation.Required),
		validation.Field(&u.LastName, validation.Required),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(requiredIf(u.password_hash == "")), validation.Length(6, 100)),
	)
}

func (i *UpdateUser) Validate() error {
	if i.UserName == nil && i.FirstName == nil && i.LastName == nil && i.Surname == nil && i.Email == nil && i.Password == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
