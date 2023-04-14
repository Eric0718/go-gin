package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	//deal with all get request
	r.POST("/api-get", apiGet)
	//deal with all post request
	r.POST("/api-post", apiPost)
}
