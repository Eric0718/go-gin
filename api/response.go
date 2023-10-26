package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(g *gin.Context, code int, data interface{}) {
	g.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  getMsg(code),
		"data": data,
	})
}

func getMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
