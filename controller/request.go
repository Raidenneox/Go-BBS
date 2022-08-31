package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var CtxUserIDKey = "userID"

var ErrorUserNotLogin = errors.New("用户未登录")

// GetCurrentUSerID GetCurrentUSer 获取当前登录的用户ID
func GetCurrentUSerID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
