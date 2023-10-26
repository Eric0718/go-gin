package routers

import (
	"go-gin/api"
	"go-gin/middle"

	"github.com/gin-gonic/gin"
)

//需要走拦截器的接口
func initApiGroup(r *gin.Engine) {
	apiGroup := r.Group("/api")
	apiGroup.Use(middle.Prepare())

	apiGroup.POST("/test", api.TestApi)
}
