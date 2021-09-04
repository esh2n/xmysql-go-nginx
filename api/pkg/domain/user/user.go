package user

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
