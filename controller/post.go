package controller

import (
	"web_app/logic"
	"web_app/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func CreatePostHandler(c *gin.Context) {
	//1.获取参数及参数的校验
	//定义模型
	p := new(models.Post)

	//validator-->binding tag
	if err := c.ShouldBind(p); err != nil {
		zap.L().Error("c.ShouldBind(p) failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//从c中拿到当前发请求的用户的ID
	userID, err := GetCurrentUSerID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	//2.创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	//3.返回响应
	ResponseSuccess(c, nil)
}
