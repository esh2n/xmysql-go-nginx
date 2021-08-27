package controller

import (
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.RouterGroup) {
	users := r.Group("/user")
	{
		u := UserController{}
		users.POST("/create", u.CreateUser)
		users.GET("/get", u.GetUser)
		users.PATCH("/update", u.UpdateUser)
		//
		// users.GET("", u.Index)
		// users.GET("/:id", u.GetUser)
		// users.POST("", u.CreateUser)
		// users.PATCH("/:id", u.UpdateUser)
		// users.DELETE("/:id", u.DeleteUser)
	}
}
