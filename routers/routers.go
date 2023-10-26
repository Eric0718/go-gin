package routers

import (
	"go-gin/api"
	"go-gin/middle"

	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	//deal with all post request
	apiGroup := r.Group("/api")
	apiGroup.Use(middle.Prepare())
	apiGroup.POST("/test", api.TestApi)

	r.GET("/test2", api.TestApi2)
}
