package models

import (
	"github.com/kamva/mgm/v3"
)

type User struct {
	mgm.DefaultModel `bson:",inline" valid:"-"`
	FullName         string `json:"fullName" bson:"fullName" valid:"required"`
	Email            string `json:"email" bson:"email" valid:"required"`
	Password         string `json:"password" bson:"password" valid:"required"`
	PhoneNumber      string `json:"phoneNumber" bson:"phoneNumber" valid:"required"`
	Level            string `json:"level" bson:"level" valid:"-"`
}

type Users struct {
	Users []User `json:"users"`
}

func NewUser(user User) *User {

	return &User{
		FullName:    user.FullName,
		Email:       user.Email,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
		Level:       user.Level,
	}
}
