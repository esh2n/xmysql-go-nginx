package model

import (
	"github.com/esh2n/xmysql-go-nginx/api/pkg/auth"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Token string
}

func NewUser(name string, id int) *User {
	token := auth.CreateTokenString(id, name)

	return &User{
		Name:  name,
		Token: string(token)}
}

func CreateUser(db *gorm.DB, user *User) (*User, error) {
	result := db.Create(&user)
	return user, result.Error
}

func GetUserByToken(db *gorm.DB, TOKEN string) (*User, error) {
	user := User{}
	result := db.Where("token", TOKEN).First(&user)

	return &user, result.Error
}

func GetLastUser(db *gorm.DB) (*User, error) {
	user := User{}
	result := db.Last(&user)

	return &user, result.Error
}

func (user *User) Update(db *gorm.DB, param map[string]interface{}) (*User, error) {
	result := db.Model(&user).Updates(param)

	return user, result.Error
}
