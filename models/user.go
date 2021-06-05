package models

import (
	"github.com/kamva/mgm/v3"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	FullName         string `json:"fullName" bson:"fullName"`
	Email            string `json:"email" bson:"email"`
	Password         string `json:"password" bson:"password"`
	PhoneNumber      string `json:"phoneNumber" bson:"phoneNumber"`
	Level            string `json:"level" bson:"level"`
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
