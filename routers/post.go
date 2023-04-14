package routers

import (
	"go-gin/logger"
	"go-gin/routers/errmsg"
	res "go-gin/routers/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func apiPost(c *gin.Context) {
	logger.PrintHttpRequest(c)
	resG := res.Gin{C: c}
	id := -1
	if Id := c.Query("id"); Id != "" {
		id = com.StrTo(Id).MustInt()
	} else {
		logger.Error("request method id[%v] empty!", Id)
		resG.Response(http.StatusBadRequest, errmsg.INVALID_PARAMS, nil)
		return
	}

	var param string
	if Param := c.Query("param"); Param == "" {
		logger.Error("request param [%v] empty!", Param)
		resG.Response(http.StatusBadRequest, errmsg.INVALID_PARAMS, nil)
		return
	} else {
		param = Param
	}

	if handler, ok := handlers[id]; ok {
		err, _ := handler(param)
		if err != nil {
			logger.Error(err)
			resG.Response(http.StatusBadRequest, errmsg.INVALID_PARAMS, nil)
			return
		}
	} else {
		logger.Error("request method id[%v] not exist!", id)
		resG.Response(http.StatusBadRequest, errmsg.INVALID_PARAMS, nil)
		return
	}

	resG.Response(http.StatusOK, errmsg.SUCCESS, nil)
	return
}
