package routes

import (
	"openim/internal/handlers"
	"openim/internal/middleware"
	"openim/internal/services"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	corsMiddleWare := cors.Default()

	// 通用中间件
	r.Use(corsMiddleWare)
	r.Use(middleware.RequestIDMiddleware())

	r.GET("/ping", pong)

	api := r.Group("/api")
	api.POST("/register", services.UserRegister) // 用户注册
	// todo 检查邮箱是否重复的接口
	api.POST("/login", handlers.LoginHandler) // 用户登录接口

	// 需要验证token的
	authApi := api.Use(middleware.AuthMiddleware())
	authApi.POST("/refreshToken", services.RefreshTokenHandler)
	authApi.GET("/userinfo", handlers.UserInfoHandler) // 需要通过token验证

	// group
	r.GET("/groups", services.OpenGroups)           // 查看所有开放的组列表
	r.POST("/group/apply", services.ApplyJoinGroup) // 申请入群
	r.GET("/my/group/list", services.MyGroupList)   // 查看自己的组列表
	r.GET("/group/info", services.GroupInfo)        // 查看某个组的信息
	r.GET("/group/members", services.GroupMembers)
	r.POST("/group/create", services.CreateGroup)
	r.POST("/group/join", services.JoinGroup)
	r.POST("/group/quit", services.QuitGroup)

	return r
}
func LogrusLogger(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		end := time.Now()

		// 记录请求日志
		log.WithFields(logrus.Fields{
			"request_id": c.MustGet("request_id"),
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"status":     c.Writer.Status(),
			"duration":   end.Sub(start),
		}).Info("Request processed")
	}
}

// test service available
func pong(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
