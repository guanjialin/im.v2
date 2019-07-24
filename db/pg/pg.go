package pg

import (
	"sync"

	"github.com/go-pg/pg"
	"im.v2/config"
)

var db *pg.DB
var dbOnce sync.Once

func init() {
	dbOnce.Do(func() {
		db = pg.Connect(&pg.Options{
			Addr:     config.PG().Addr,
			User:     config.PG().User,
			Password: config.PG().Password,
			Database: config.PG().Database,
		})
	})
}

func DB() *pg.DB {
	return db
}
