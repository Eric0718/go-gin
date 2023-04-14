package routers

import (
	"go-gin/app/userinfo"

	"github.com/gin-gonic/gin"
)

func initUserGroup(r *gin.Engine) {
	r.POST("/dealBlacklist", userinfo.DealBlacklist)
}
