package routers

import (
	"go-gin/api"

	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	//deal with all post request
	r.POST("/api/test", api.TestApi)
}
