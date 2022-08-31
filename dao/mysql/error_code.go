package mysql

import "errors"

var (
	ErrorUserExist     = errors.New("用户已存在")
	ErrorUserNotExist  = errors.New("无此用户")
	ErrorWrongPassword = errors.New("用户名或密码错误")
	ErrorInvalidID     = errors.New("无效的ID")
)
