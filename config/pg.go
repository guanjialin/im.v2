package config

import "log"

const pgPath = "/pgdb.json"

type pg struct {
	Addr     string	`json:"addr"`
	User     string	`json:"user"`
	Password string	`json:"password"`
	Database string	`json:"database"`
}

func PG() *pg {
	config := new(pg)
	if err := parseConfig(pgPath, config); err != nil {
		log.Fatalln(err)
		return nil
	}

	return config
}
