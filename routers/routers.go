package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	initNftGroup(r)
	initUserGroup(r)
}
