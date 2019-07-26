package model

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// 用户信息表
type User struct {
	basic
	ObjectID	string	`sql:"object_id,type:varchar(32),pk"`
	Nickname	string	`sql:"nick_name, notnull"`
	Password	string	`sql:"password, type:varchar(32), notnull"`
	Portrait	string	`sql:"portrait"`
}

func (u User) CreateTable(tx *pg.Tx) error {
	return tx.CreateTable(&u, &orm.CreateTableOptions{
		IfNotExists:true,
	})
}

func (u User) CreateIndex(tx *pg.Tx) error {

	return nil
}


func (u User) CreateTrigger(tx *pg.Tx) error {
	return nil
}