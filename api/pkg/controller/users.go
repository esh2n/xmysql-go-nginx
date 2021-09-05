package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	model "github.com/esh2n/xmysql-go-nginx/api/pkg/domain/user"
	infra "github.com/esh2n/xmysql-go-nginx/api/pkg/infra/user"
)

type UserController struct{}

type UserParam struct {
	Name string `json:"name" binding:"required,min=1,max=50"`
}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var param UserParam
	if err := c.BindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ur := infra.NewUserRepository()

	user := model.NewUser(param.Name)
	newUser, err := ur.CreateUserInTransaction(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed on CreateUser()"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": newUser})
}

func (uc *UserController) GetUser(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	ur := infra.NewUserRepository()

	user, err := ur.GetUserByToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	ur := infra.NewUserRepository()

	user, err := ur.GetUserByToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	var param UserParam
	if err := c.BindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateParam := map[string]interface{}{
		"name":  param.Name,
		"token": token,
	}

	_, err = ur.UpdateUser(user, updateParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user update failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
