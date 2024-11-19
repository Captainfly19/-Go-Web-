package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	//将列定义为唯一键，就可以有效避免用户名重复的问题
	Username string `gorm:"unique"`
	Password string
}
