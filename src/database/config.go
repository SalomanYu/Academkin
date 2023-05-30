package database

import (
	"database/sql"
	"fmt"
	"strings"
)

type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	db       *sql.DB
}

func (d *Database) InitDatabase() (err error) {
	dbUrl := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", d.Host, d.Port, d.User, d.Password, d.Name)
	d.db, err = sql.Open("postgres", dbUrl)
	return
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func replaceAllQuotes(text string) string {
	text = strings.ReplaceAll(text, `"`, "`")
	text = strings.ReplaceAll(text, `'`, "`")
	return text
}