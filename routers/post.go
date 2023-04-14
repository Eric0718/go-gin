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
		logger.Errorf("request id[%v] empty!", Id)
		resG.Response(http.StatusBadRequest, errmsg.INVALID_PARAMS, nil)
		return
	}
	var method string
	if method = c.Query("method"); method == "" {
		logger.Errorf("request method[%v] empty!", method)
		resG.Response(http.StatusBadRequest, errmsg.INVALID_PARAMS, nil)
		return
	}

	var param string
	if param = c.Query("param"); param == "" {
		logger.Errorf("request param [%v] empty!", param)
		resG.Response(http.StatusBadRequest, errmsg.INVALID_PARAMS, nil)
		return
	}

	if f, ok := handlers[id]; ok {
		if idToFunc[id] != method {
			logger.Errorf("request method[%v] not registed!", method)
			resG.Response(http.StatusBadRequest, errmsg.INVALID_PARAMS, nil)
			return
		}
		err, _ := f(param)
		if err != nil {
			logger.Error(err)
			resG.Response(http.StatusBadRequest, errmsg.INVALID_PARAMS, nil)
			return
		}

	} else {
		logger.Errorf("request id[%v] not exist!", id)
		resG.Response(http.StatusBadRequest, errmsg.INVALID_PARAMS, nil)
		return
	}

	resG.Response(http.StatusOK, errmsg.SUCCESS, nil)
	return
}
