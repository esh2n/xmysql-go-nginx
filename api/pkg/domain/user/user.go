package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    uint   `gorm:"primary_key"`
	Name  string `json:"name"`
	Token string
}

func NewUser(name string) *User {
	return &User{
		Name: name,
	}
}
