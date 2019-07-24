package model

import (
	"time"
	"github.com/jinzhu/gorm"
)

type basic struct {
	ID       int64     `sql:"id, notnull"`
	CreateAt time.Time `sql:"create_at, notnull, type:timestamp with timezone, default:now()"`
	UpdateAt time.Time `sql:"update_at, notnull, type:timestamp with timezone, default:now()"`
	DeleteAt time.Time `sql:"delete_at, type:timestamp with timezone"`
}

func Register() {

}
