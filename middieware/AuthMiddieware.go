package middieware

import (
	"OceanLearn/common"
	"OceanLearn/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

/**
	Token解析与校验-中间件
 */
func AuthMiddieware() gin.HandlerFunc  {

	return func(context * gin.Context){
		// 从请求中获取authorization header,
		tokenString := context.GetHeader("Authorization")

		// 校验token 格式是否正确
		if tokenString =="" || !strings.HasPrefix(tokenString,"Bearer"){
			// 验证token不通过，返回错误信息，请求终止
			context.JSON(http.StatusUnauthorized,gin.H{"":401,"msg":"权限不足"})
			context.Abort()
			return
		}
		// 字符串截取，从第7位开截取到最后
		tokenString = tokenString[7:]

		// 解析token
		token,claims,err := common.ParesToken(tokenString)
		if(err != nil || !token.Valid){
			// 解析token失败，或者，请求终止
			context.JSON(http.StatusUnauthorized,gin.H{"":401,"msg":"权限不足"})
			context.Abort()
			return
		}
		// token验证通过之后，从claim中获取userId
		userId := claims.UserId

		// 从数据库中获取查询userID对应的用户信息
		DB := common.GetMysqlDB()
		var user model.User
		DB.First(&user,userId)

		if user.ID ==0 {
			// 如果用户信息查不到，则代表该用户不存在或权限不足
			context.JSON(http.StatusUnauthorized,gin.H{"":401,"msg":"权限不足"})
			context.Abort()
			return
		}
		// 如果用户信息存在，则将用户信息存在请求中
		context.Set("user",user)
		// 让请求走下去
		context.Next()
	}
}
