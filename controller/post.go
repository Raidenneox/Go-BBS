package controller

import (
	"strconv"
	"web_app/logic"
	"web_app/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// CreatePostHandler 创建帖子
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

// GetPostDetailHandler 获取帖子详情的处理函数
func GetPostDetailHandler(c *gin.Context) {
	// 1. 获取参数（从URL中获取帖子的id）
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 2. 根据id取出帖子数据（查数据库）
	data, err := logic.GetPostByID(pid)
	if err != nil {
		zap.L().Error("logic.GetPostById(pid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, data)
}

func GetPostListHandler(c *gin.Context) {

	page, size := GetPageInfo(c)

	//查询到所有的帖子并以列表的形式返回
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
	}
	//返回目前的参数列表
	ResponseSuccess(c, data)
}

// GetPostListHandler2 根据前端传来参数动态获取帖子的列表
//按创建时间/分数排序

//1.获取请求的query string参数
//2.去redis查询ID值
//3.根据ID去数据库查询帖子详细信息

func GetPostListHandler2(c *gin.Context) {
	//GET请求参数(querystring): /api/v1/post2?p=1&s=10&o=time

	//初始化结构体时指定初始参数
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}
	if err := c.ShouldBind(p); err != nil {
		zap.L().Error("GetPostListHandler2() with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	data, err := logic.GetPostListNew(p) //更新,两个接口合二为一
	//获取数据

	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
	}
	//返回目前的参数列表
	ResponseSuccess(c, data)
}

//根据社区去查询帖子列表
//
//func GetCommunityPostListHandler(c *gin.Context) {
//
//	//初始化结构体时指定初始参数
//	p := &models.ParamCommunityPostList{
//		ParamPostList: models.ParamPostList{
//			Page:  1,
//			Size:  10,
//			Order: models.OrderTime,
//		},
//	}
//	if err := c.ShouldBind(p); err != nil {
//		zap.L().Error("GetPostListHandler2() with invilid param", zap.Error(err))
//		ResponseError(c, CodeInvalidParam)
//		return
//	}
//	//获取数据
//
//	if err != nil {
//		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
//		ResponseError(c, CodeServerBusy)
//		return
//	}
//	//返回目前的参数列表
//	ResponseSuccess(c, data)
//
//}
