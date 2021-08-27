package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fuhiz/docker-go-sample/app/pkg/connecter"
	"github.com/fuhiz/docker-go-sample/app/pkg/controller"
)

func main() {
	connecter.Setup()
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Response OK")
	})

	r := router.Group("/")
	controller.Setup(r)

	router.Run(":3000")
}
