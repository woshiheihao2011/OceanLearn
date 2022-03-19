package model

import "github.com/jinzhu/gorm"

/**
声明user对象，并且指定表结构字段名、字符长度、是否为null
*/
type User struct {
	gorm.Model
	Name string `gorm:"varchar(20);not null"`
	Telephone string `gorm:"varchar(11);not null"`
	Password string  `gorm:"varchar(255);not null"`
}
