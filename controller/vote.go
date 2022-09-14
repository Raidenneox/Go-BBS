package controller

import (
	"web_app/logic"
	"web_app/models"

	"go.uber.org/zap"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

func PostVoteHandler(c *gin.Context) {

	//参数校验
	p := new(models.ParamVoteData)
	if err := c.ShouldBind(p); err != nil {
		err, ok := err.(validator.ValidationErrors) //类型断言
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(err.Translate(trans)) //翻译并取出错误中的结构体表示
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}
	//获取当前请求的用户ID
	userID, err := GetCurrentUSerID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
	}
	err = logic.VoteForPost(userID, p)
	if err != nil {
		zap.L().Error("logic.VoteForPost()failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
