package pgdb

import (
	"github.com/sirupsen/logrus"
	"im.v2/config"
	"sync"
	"time"

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
		logrus.Debugf("Postgres connected: %s:%s@%s/%s", config.PG().User,
			config.PG().Password, config.PG().Addr, config.PG().Database)

		db.AddQueryHook(&dbLog{})
	})

	return db
}

type dbLog struct {
	time time.Time
}

func (d *dbLog) BeforeQuery(event *pg.QueryEvent) {
	d.time = time.Now().UTC()
}

func (d dbLog) AfterQuery(event *pg.QueryEvent) {
	sql, err := event.FormattedQuery()
	if err != nil {
		logrus.Error("SQL: ", sql)
	}
	logrus.Debugf("SQL: %s %s\n", time.Since(d.time), sql)
}