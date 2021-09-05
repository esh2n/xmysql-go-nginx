package user

import (
	"github.com/esh2n/xmysql-go-nginx/api/pkg/auth"
	u "github.com/esh2n/xmysql-go-nginx/api/pkg/domain/user"
	"github.com/esh2n/xmysql-go-nginx/api/pkg/infra/connecter"
)

type UserRepository struct {
	c *connecter.Connecter
}

func NewUserRepository() *UserRepository {
	c := connecter.NewConnecter()
	return &UserRepository{
		c: c,
	}
}

func (ur *UserRepository) CreateUser(user *u.User) (*u.User, error) {
	result := ur.c.DB.Create(&user)
	return user, result.Error
}

func (ur *UserRepository) UpdateToken(user *u.User) (*u.User, error) {
	token := auth.CreateTokenString(user.ID, user.Name)
	return ur.UpdateUser(user, map[string]interface{}{"token": token})
}

func (ur *UserRepository) CreateUserInTransaction(user *u.User) (*u.User, error) {
	data, err := ur.c.TransactinHandler(func() (interface{}, error) {
		newUser, err := ur.CreateUser(user)
		if err != nil {
			return nil, err
		}
		updatedUser, errEvent := ur.UpdateToken(newUser)
		if errEvent != nil {
			return nil, errEvent
		}
		return updatedUser, nil
	})
	if err != nil {
		return nil, err
	}
	return data.(*u.User), nil
}

func (ur *UserRepository) GetUserByToken(TOKEN string) (*u.User, error) {
	user := u.User{}
	result := ur.c.DB.Where("token", TOKEN).First(&user)
	return &user, result.Error
}

func (ur *UserRepository) GetLastUser() (*u.User, error) {
	user := u.User{}
	result := ur.c.DB.Last(&user)
	return &user, result.Error
}

func (ur *UserRepository) UpdateUser(user *u.User, param map[string]interface{}) (*u.User, error) {
	result := ur.c.DB.Model(&user).Updates(param)
	return user, result.Error
}
