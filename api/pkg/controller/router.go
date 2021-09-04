package controller

import (
	"github.com/gin-gonic/gin"
)

func Setup(rg *gin.RouterGroup) {
	users := rg.Group("/user")
	{
		uc := NewUserController()
		users.POST("/create", uc.CreateUser)
		users.GET("/get", uc.GetUser)
		users.PATCH("/update", uc.UpdateUser)
	}
}
