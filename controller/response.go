package controller

import "C"
import "github.com/gin-gonic/gin"

// ResponseData
//{
//	"code":1001,//程序中的错误码
//	"msg":xx,	//提示信息
//	"data":{},	//数据
//}
type ResponseData struct {
	Code ResCode `json:"code"`
	Msg  any     `json:"msg"`
	Data any     `json:"data,omitempty"`
}

func ResponseSuccess(c *gin.Context, data any) {
	rd := &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	}
	c.JSON(200, rd)
}
func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(200, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg any) {
	c.JSON(200, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
