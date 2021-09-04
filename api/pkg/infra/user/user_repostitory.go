package user

import (
	u "github.com/esh2n/xmysql-go-nginx/api/pkg/domain/user"
	"github.com/esh2n/xmysql-go-nginx/api/pkg/infra/connecter"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepository {
	c := connecter.NewConnecter()
	return &UserRepository{
		DB: c.DB,
	}
}

func (ur *UserRepository) CreateUser(user *u.User) (*u.User, error) {
	result := ur.DB.Create(&user)
	return user, result.Error
}

func (ur *UserRepository) GetUserByToken(TOKEN string) (*u.User, error) {
	user := u.User{}
	result := ur.DB.Where("token", TOKEN).First(&user)
	return &user, result.Error
}

func (ur *UserRepository) GetLastUser() (*u.User, error) {
	user := u.User{}
	result := ur.DB.Last(&user)
	return &user, result.Error
}

func (ur *UserRepository) UpdateUser(user *u.User, param map[string]interface{}) (*u.User, error) {
	result := ur.DB.Model(&user).Updates(param)
	return user, result.Error
}
