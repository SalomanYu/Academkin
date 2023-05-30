package database

import (
	"database/sql"
	"fmt"
)

type Database struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
	db       *sql.DB
}

func (d *Database) InitDatabase() (err error) {
	dbUrl := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", d.host, d.port, d.user, d.password, d.dbname)
	d.db, err = sql.Open("postgres", dbUrl)
	return
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
