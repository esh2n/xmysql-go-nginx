package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/esh2n/xmysql-go-nginx/api/pkg/connecter"
	"github.com/esh2n/xmysql-go-nginx/api/pkg/controller"
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
