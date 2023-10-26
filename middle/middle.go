package middle

import (
	"go-gin/logger"

	"github.com/gin-gonic/gin"
)

func Prepare() gin.HandlerFunc {
	return func(c *gin.Context) {
		agent := c.GetHeader("User-Agent")
		logger.Debugf("test prepare===========%v", agent)
	}
}
