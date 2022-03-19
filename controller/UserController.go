package controller

import (
	"OceanLearn/common"
	"OceanLearn/dto"
	"OceanLearn/model"
	"OceanLearn/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

/**
	注册用户
 */
func Register(context *gin.Context){
	// 获取参数,PostMan测试传参时，必须在Form-Data里录入参数
	//name := context.PostForm("name")
	//telephone := context.PostForm("telephone")
	//password := context.PostForm("password")

	var user model.User
	// 从请求中获取参数，并绑定至user对象中
	if err := context.ShouldBind(&user); err != nil{
		context.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":"422",
			"msg":"入参数据格式/参数名不正确！",
		})
		return
	}
	name := user.Name
	telephone := user.Telephone
	password := user.Password

	// 数据验证1:电话号码长度不等于11位时，流程终止
	if len(telephone) != 11{
		// 写法1：
		context.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":"422",
			"msg":"手机号必须为11位！",
		})
		// 写法2：
		//context.JSON(422,map[String] interface{}{
		//	"code":"422",
		//	"msg":"手机号必须为11位！",
		//})
		return
	}

	// 数据验证2：password 长度不能少于6位
	if len(password) < 6 {
		context.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":"422",
			"msg":"password必须大于6位！",
		})
		return
	}

	// 数据校验3：如果name值为null，则默认赋值一个10位的随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	log.Println(name,telephone,password)

	db := common.GetMysqlDB()
	// 判断手机号是否存在
	if isTelephoneExist(db,telephone){
		context.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":"422",
			"msg":"手机号码已经存在！",
		})
		return
	}

	// 对用户密码进行加密
	hasePassword,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{
			"code":"422",
			"msg":"密码加密失败！",
		})
		return
	}

	// 创建用户
	var newUser = model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasePassword),
	}
	db.Create(&newUser)

	// 返回结果
	context.JSON(200,gin.H{
		"code":"200",
		"message":"注册成功",
	})
}

/**
查询数据库，看电话号码是否在数据库中存在
存在则返回true
不存在则返回 false
*/
func isTelephoneExist(db *gorm.DB,telephone string) bool{
	var userObject model.User
	// 传入查询条件，并返回第一条数据给userObject
	db.Where("telephone = ?",telephone).First(&userObject)
	if userObject.ID != 0 {
		return true
	}
	return false
}

/**
	模拟登录
 */
func Login(context *gin.Context){
	// 获取数据连接
	DB := common.GetMysqlDB()
	var user model.User
	// 从请求中获取参数，并绑定至user对象中
	if err := context.ShouldBind(&user); err != nil{
		context.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":"422",
			"msg":"入参数据格式/参数名不正确！",
		})
		return
	}
	name := user.Name
	telephone := user.Telephone
	password := user.Password


	// 数据验证1:电话号码长度不等于11位时，流程终止
	if len(telephone) != 11{
		// 写法1：
		context.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":"422",
			"msg":"手机号必须为11位！",
		})
		return
	}

	// 数据验证2：password 长度不能少于6位
	if len(password) < 6 {
		context.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":"422",
			"msg":"password必须大于6位！",
		})
		return
	}
	// 判断手机是否存在
	DB.Where("telephone =? ",telephone).First(&user)
	if user.ID == 0{
		context.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":"422",
			"msg":"手机号码不存在",
		})
		return
	}
	// 判断密码是否正确,报文传过来的密码与数据库查询出的密码做比较
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password)); err != nil{
		context.JSON(http.StatusBadRequest,gin.H{
			"code":"400",
			"msg":"密码错误",
		})
		return
	}
	// 发放token
	token ,err := common.ReleaseToken(user)

	if err !=nil{
		context.JSON(http.StatusInternalServerError,gin.H{
			 "code":"500",
			 "msg":"系统异常",
		})
		return
	}
	// 加密算法：HS256
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjQsImV4cCI6MTY0NzgzNDQwMiwiaWF0IjoxNjQ3MjI5NjAyLCJpc3MiOiJvY2VhbmxlYXJuLnRlY2giLCJzdWIiOiJ1c2VyIHRva2VuIn0.Akba7DByWS4_8tjfzhUMCzo3zuRAf8zvwh1iLFaEBwE
	// token格式：headerBase64.claimsBase64.headerBase64+claimsBase64+jwtKey
	// headerBase64 代表头信息
	// claimsBase64 代表token对象初始化的body信息
	// headerBase64+claimsBase64+jwtKey 代表 头+body+公钥生成的Hash值
	// 返回结果
	context.JSON(200,gin.H{
		"code":"200",
		"data":gin.H{
			"token":token,
		},
		"message":"登录成功:"+name,
	})
}
/**
	获取用户token信息，方便鉴权校验
 */
func Info(context *gin.Context){
	user,_:= context.Get("user")
	context.JSON(http.StatusOK,gin.H{"code":"200","data":gin.H{"user":dto.ToUserDto(user.(model.User))}})
}