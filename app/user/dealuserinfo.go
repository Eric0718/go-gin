package user

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go-gin/dbutil"
	"go-gin/logger"
	"go-gin/models"
)

func Blacklist(param string) (error, interface{}) {
	pm, err := hex.DecodeString(param)
	if err != nil {
		logger.Error("err:", err)
		return err, nil
	}

	var ureq models.UserRequest
	err = json.Unmarshal(pm, &ureq)
	if err != nil {
		logger.Error("err:", err)
		return err, nil
	}

	email := ureq.Email
	state := ureq.Flag

	if !(state == 0 || state == 1 || state == 3) {
		return fmt.Errorf("state[%v] not wright!", state), nil
	}
	logger.Debug("email:", email, "state:", state)
	var user models.UserBase
	ok, err := dbutil.SqlEngine.Where("email = ?", email).Get(&user)
	if !ok || err != nil {
		logger.Error("err:", err)
		return fmt.Errorf("email[%v] not exist or sql error!", email), nil
	}

	_, err = dbutil.SqlEngine.Where("email = ?", email).Cols("del_flag").Update(models.UserBase{DelFlag: state})
	if err != nil {
		logger.Error("err:", err)
		return err, nil
	}
	return nil, nil
}
