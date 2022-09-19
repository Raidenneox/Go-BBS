package router

import (
	"time"
	"web_app/controller"
	"web_app/logger"
	"web_app/middlewares"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	v1 := r.Group("/v1/api")

	v1.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	//登陆的处理函数
	v1.POST("/signup", controller.SignUpHandler) //注册的处理函数
	v1.POST("/login", controller.LoginHandler)
	//令牌桶算法
	//10个令牌,每两秒钟添加一个

	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.GET("/posts", controller.GetPostListHandler)

		//根据时间或分数获取帖子列表
		v1.GET("/posts2", controller.GetPostListHandler2)
	}
	v1.Use(middlewares.JWTAuthMiddleware(), middlewares.RateLimitMiddleware(2*time.Second, 10)) //应用JWT认证中间件
	{
		v1.POST("/post", controller.CreatePostHandler)
		v1.POST("/vote", controller.PostVoteHandler)

	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "404 NOT FOUND",
		})
	})
	return r
}
