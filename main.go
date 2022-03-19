package main

import (
	"OceanLearn/common"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)



func main(){
	// 初始化数据库连接对象
	db := common.GetMysqlDB()
	// 这句代码什么意思，老子没看懂？？？？？
	defer db.Close()

	//  获取默认路由引擎
	r := gin.Default()

	// 设置get请求，并返回消息
	//r.GET("/ping",func(c *gin.Context){
	//	c.JSON(200,gin.H{
	//		"message":"pong",
	//	})
	//})
	// 初始化请求API
	r = CollectRoute(r)
	panic(r.Run(":8081"))
}
