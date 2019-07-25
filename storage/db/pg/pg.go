package pg

import (
	"im.v2/config"
	"sync"

	"github.com/go-pg/pg"
)

var db *pg.DB
var dbOnce sync.Once

func DB() *pg.DB {
	dbOnce.Do(func() {
		db = pg.Connect(&pg.Options{
			Addr:     config.PG().Addr,
			User:     config.PG().User,
			Password: config.PG().Password,
			Database: config.PG().Database,
		})
	})

	return db
}
