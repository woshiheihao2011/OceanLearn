package common

import (
	"OceanLearn/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
/**
初始化数据库连接配置,并返回DB对象
*/
func initDB() *gorm.DB{
	driverName := "mysql"
	host :="127.0.0.1"
	port :="3306"
	dataBase := "ginessential"
	userName := "root"
	password := "88888888"
	charset := "utf8mb4"

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		userName,
		password,
		host,
		port,
		dataBase,
		charset)
	fmt.Println("driverName :"+args)
	db, err := gorm.Open(driverName,args)
	if err != nil{
		panic("failed to connect database, err : "+err.Error())
	}
	// 根据对象自动创建表
	db.AutoMigrate(&model.User{})
	return db
}

/**
	获取数据库链接对象
 */
func GetMysqlDB() *gorm.DB{
	DB = initDB()

	return DB
}