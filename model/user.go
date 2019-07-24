package model

import "github.com/jinzhu/gorm"

// 用户信息表
type User struct {
	gorm.Model
	ObjectID	string	`type:varchar(32)`
	Nickname	string
	Password	string
}