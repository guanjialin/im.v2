package model

import (
	"github.com/go-pg/pg/orm"
	"im.v2/storage/db/pg"
)

// 用户信息表
type User struct {
	basic
	ObjectID	string	`sql:"object_id, type:varchar(32), unique_index, not null"`
	Nickname	string	`sql:"nick_name, index, not null"`
	Password	string	`sql:"password, type:varchar(32), not null"`
	Portrait	string	`sql:"portrait"`
}

func (u User) CreateTable() error {
	return pg.DB().CreateTable(&u, &orm.CreateTableOptions{
		IfNotExists:true,
	})
}