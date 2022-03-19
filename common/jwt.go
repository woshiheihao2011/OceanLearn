package common

import (
	"OceanLearn/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 声明jwtKey变量
var jwtKey =[]byte("a_secret_crect")

/**
	声明一个Model对象，必须要引入jwt组件，封装token对象
 */
type Claims struct {
	UserId uint
	jwt.StandardClaims
}


/**
	通过User对象生成 7*24小时token
 */
func ReleaseToken(user model.User) (string,error){
	// 设置token获取的时间
	expirationTime := time.Now().Add( 7 * 24 * time.Hour)

	// 设置并初始化 token
	claims := &Claims{
		UserId:  user.ID,
		StandardClaims :jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), //初始化token有效时间
			IssuedAt: time.Now().Unix(),//初始化token发放时间
			Issuer: "oceanlearn.tech",//初始化发放的token
			Subject: "user token",//初始化主题
		},
	}
	// 基于claims（用户基本信息）使用SigningMethodHS256加密方式初始化,并返回一个token指针
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	// 生成token
	tokenString, err:= token.SignedString(jwtKey)
	// 异常校验
	if err != nil{
		return "",err
	}
	return tokenString,nil
}

/**
	解析token
 */
func ParesToken(tokenString string) (*jwt.Token, *Claims,error){

	// 初始化Claims
	claims := &Claims{}
	// 解析token,解析成功返回 jwtKye（公钥）
	token, err := jwt.ParseWithClaims(tokenString,claims, func(token  *jwt.Token) (interface{}, error) {

		return jwtKey, nil
	})

	return token,claims,err
}