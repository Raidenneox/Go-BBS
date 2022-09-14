package router

import (
	"web_app/controller"
	"web_app/logger"
	"web_app/middlewares"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	v1 := r.Group("/v1/api")

	//登陆的处理函数
	v1.POST("/signup", controller.SignUpHandler) //注册的处理函数
	v1.POST("/login", controller.LoginHandler)
	v1.Use(middlewares.JWTAuthMiddleware()) //应用JWT认证中间件

	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)
	}
	{
		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.GET("/posts", controller.GetPostListHandler)
	}

	v1.POST("/vote", controller.PostVoteHandler)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "404 NOT FOUND",
		})
	})
	return r
}
