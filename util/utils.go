package util

import (
	"math/rand"
	"time"
)

/**
生成随机字符串并返回，length字符长度， 返回字符串
*/
func RandomString(length int) string{
	// 声明秘钥
	var letters = []byte("alksjdlkfjaskdjasjdhfajsdfjasdfasdfasdf")

	result  :=  make([]byte,length)

	rand.Seed(time.Now().Unix())

	for i:= range  result{
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}
