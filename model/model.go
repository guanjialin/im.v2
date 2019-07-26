package model

import (
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"im.v2/config"
	"im.v2/storage/db/pgdb"
)

type basicInterface interface {
	CreateTable(tx *pg.Tx) error
	CreateIndex(tx *pg.Tx) error
	CreateTrigger(tx *pg.Tx) error
}

type basic struct {
	ID        int64     `sql:"id, pk"`
	CreateAt  time.Time `sql:"create_at,notnull, type:timestamp with time zone, default:now()"`
	UpdateAt  time.Time `sql:"update_at,notnull, type:timestamp with time zone, default:now()"`
	DeletedAt time.Time `sql:"delete_at,type:timestamp with time zone" pg:",soft_delete"`
}

func init() {
	orm.SetTableNameInflector(func(s string) string {
		return config.PG().Schema + "." + s
	})
}

func InitTable() error {
	return pgdb.DB().RunInTransaction(func(tx *pg.Tx) error {
		for _, v := range models {
			err := v.CreateTable(tx)
			if err != nil {
				return err
			}

			err = v.CreateTrigger(tx)
			if err != nil {
				return err
			}
		}

		return nil
	})
}
