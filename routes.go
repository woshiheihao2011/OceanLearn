package main

import (
	"OceanLearn/controller"
	"OceanLearn/middieware"
	"github.com/gin-gonic/gin"
)

/**
	API声明管理类
 */
func CollectRoute(r *gin.Engine) *gin.Engine{
	// 通过Use方法，注入跨域请求配置的中间件, （Use方法：中间件注入）
	r.Use(middieware.CORSMiddleware())
	// 设置访问API
	r.POST("/api/auth/register",controller.Register)
	r.POST("/api/auth/login",controller.Login)
	r.GET("/api/auth/info",middieware.AuthMiddieware(),controller.Info)
	return r
}
