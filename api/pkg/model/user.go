package model

import (
	"gorm.io/gorm"
)

// func CreateUser(db *gorm.DB, user *User) (*User, error) {
// 	result := db.Create(&user)

// 	return user, result.Error
// }

// func GetUsers(db *gorm.DB) ([]*User, error) {
// 	users := []*User{}
// 	result := db.Find(&users)

// 	return users, result.Error
// }

// func GetUserById(db *gorm.DB, ID int) (*User, error) {
// 	user := User{}
// 	result := db.First(&user, ID)

// 	return &user, result.Error
// }

// func (user *User) Delete(db *gorm.DB) (*User, error) {
// 	result := db.Delete(&user)

// 	return user, result.Error
// }

//

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Token string `json:"token"`
}

func NewUser(name string, token string) *User {
	return &User{
		Name:  name,
		Token: token}
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

func (user *User) Update(db *gorm.DB, param map[string]interface{}) (*User, error) {
	result := db.Model(&user).Updates(param)

	return user, result.Error
}
