package routers

import (
	"go-gin/api"

	"github.com/gin-gonic/gin"
)

//不需要走拦截器的接口
func initPublicGroup(r *gin.Engine) {
	r.GET("/test2", api.TestApi2)
}
