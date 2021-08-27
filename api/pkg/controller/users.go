package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/esh2n/xmysql-go-nginx/api/pkg/connecter"
	"github.com/esh2n/xmysql-go-nginx/api/pkg/model"
)

type UserController struct{}

type UserParam struct {
	Name string `json:"name" binding:"required,min=1,max=50"`
}

func (self *UserController) CreateUser(c *gin.Context) {
	var param UserParam
	if err := c.BindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lastUser, err := model.GetLastUser(connecter.DB())
	lastID := 0
	if err == nil {
		lastID = int(lastUser.ID)
	}

	newUser := model.NewUser(param.Name, lastID+1)
	user, err := model.CreateUser(connecter.DB(), newUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed on CreateUser()"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (self *UserController) GetUser(c *gin.Context) {
	token := c.Request.Header.Get("x-token")

	user, err := model.GetUserByToken(connecter.DB(), token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (self *UserController) UpdateUser(c *gin.Context) {
	token := c.Request.Header.Get("x-token")

	user, err := model.GetUserByToken(connecter.DB(), token)
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

	_, err = user.Update(connecter.DB(), updateParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user update failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
