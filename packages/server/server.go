package server

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

func RunServer() {
	// https://gin-gonic.com/docs/quickstart/
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:$PORT
}
