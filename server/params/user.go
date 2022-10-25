package params

import (
	"errors"
	"golang-batch6-group3/server/model"

	"github.com/go-playground/validator/v10"
)

type UserCreate struct {
	Fullname   string `validate:"required"`
	Email      string `validate:"required"`
	Password   string `validate:"required"`
	Gender     string `validate:"required"`
	Street     string `validate:"required"`
	CityId     int    `validate:"required"`
	ProvinceId int    `validate:"required"`
}

func Validate(u interface{}) error {
	err := validator.New().Struct(u)
	if err == nil {
		return nil
	}
	myErr := err.(validator.ValidationErrors)
	errString := ""
	for _, e := range myErr {
		errString += e.Field() + " is " + e.Tag()
	}
	return errors.New(errString)
}

type UserUpdate struct {
	Email    string
	Password string
}
type UserLogin struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

func (u *UserCreate) ParseToModel() *model.User {
	return &model.User{
		Fullname:   u.Fullname,
		Email:      u.Email,
		Password:   u.Password,
		Gender:     u.Gender,
		Street:     u.Street,
		CityId:     u.CityId,
		ProvinceId: u.ProvinceId,
	}
}
