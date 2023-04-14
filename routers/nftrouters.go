package routers

import (
	"github.com/gin-gonic/gin"
)

func initNftGroup(r *gin.Engine) {
	r.GET("/testGin", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Gin!",
		})
	})
}
