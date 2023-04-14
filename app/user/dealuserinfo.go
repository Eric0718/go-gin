package user

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go-gin/dbutil"
	"go-gin/logger"
	"go-gin/models"
	"go-gin/util/resmsg"
)

func Blacklist(param string) (int, interface{}) {
	pm, err := hex.DecodeString(param)
	if err != nil {
		logger.Error("err:", err)
		return resmsg.INVALID_PARAMS, nil
	}

	var ureq models.UserRequest
	err = json.Unmarshal(pm, &ureq)
	if err != nil {
		logger.Error("err:", err)
		return resmsg.INVALID_PARAMS, nil
	}

	email := ureq.Email
	state := ureq.Flag

	if !(state == 0 || state == 1 || state == 3) {
		logger.Error("err:", fmt.Errorf("state[%v] not wright!", state))
		return resmsg.INVALID_PARAMS, nil
	}
	logger.Debug("email:", email, "state:", state)
	var user models.UserBase
	ok, err := dbutil.SqlEngine.Where("email = ?", email).Get(&user)
	if !ok || err != nil {
		logger.Error("err:", fmt.Errorf("email[%v] not exist or sql error!", email))
		return resmsg.INVALID_PARAMS, nil
	}

	_, err = dbutil.SqlEngine.Where("email = ?", email).Cols("del_flag").Update(models.UserBase{DelFlag: state})
	if err != nil {
		logger.Error("err:", err)
		return resmsg.INVALID_PARAMS, nil
	}
	return resmsg.SUCCESS, nil
}
