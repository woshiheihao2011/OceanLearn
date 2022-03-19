package middieware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
	解决跨域问题的中间件
	主要设置请求头的跨域关键参数

 */
func CORSMiddleware() gin.HandlerFunc{

	return func(context *gin.Context){

		// 设置允许访问的-origin("http://localhost:8080")共享。(前后端分离解决跨域的重要配置)
		context.Writer.Header().Set("Access-Control-Allow-Origin","http://localhost:8080")
		// 设置允许访问的-Access-Control-Allow-Methods 和 Access-Control-Allow-Headers 的缓存时间
		context.Writer.Header().Set("Access-Control-Max-Age","86400")
		// 设置允许访问的-请求方式：POST, GET, OPTIONS，*代表全部
		context.Writer.Header().Set("Access-Control-Allow-Methods","*")
		// 设置允许访问的-请求头信息
		context.Writer.Header().Set("Access-Control-Allow-Headers","*")
		// 设置是否可以将对请求的响应暴露给页面。返回true则可以，其他值均不可以。
		context.Writer.Header().Set("Access-Control-Allow-Credentials","true")

		if context.Request.Method == http.MethodOptions{
			context.AbortWithStatus(http.StatusOK)
		}else {
			context.Next()
		}
	}
}
