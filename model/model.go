package model

import (
	"github.com/go-pg/pg"
	"log"
	"os"
	"time"
)

func init()  {
	pg.SetLogger(log.New(os.Stdout, "im", log.Llongfile))
}

var models = []basicInterface{
	User{},
}

type basicInterface interface {
	CreateTable() error
}

type basic struct {
	ID       int64     `sql:"id, notnull"`
	CreateAt time.Time `sql:"create_at, notnull, type:timestamp with timezone, default:now()"`
	UpdateAt time.Time `sql:"update_at, notnull, type:timestamp with timezone, default:now()"`
	DeleteAt time.Time `sql:"delete_at, type:timestamp with timezone"`
}

func CreateTable() error {
	for _, v := range models {
		err := v.CreateTable()
		if err != nil {
			return err
		}
	}

	return nil
}