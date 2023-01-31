package database

import (
	"github.com/SalomanYu/Academkin/src/models"
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

func Connect() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	checkErr(err)
	defer db.Close()
	err = db.Ping()
	checkErr(err)
}

func AddMultipleVuzes(vuzes []models.Vuz) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	checkErr(err)
	defer db.Close()
	valueStrings := []string{}
	valueArgs := []interface{}{}
	valueInsertCount := 1
	for _, vuz := range vuzes {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", valueInsertCount, valueInsertCount+1, valueInsertCount+2, valueInsertCount+3, valueInsertCount+4))
		valueArgs = append(valueArgs, vuz.ShortName)
		valueArgs = append(valueArgs, vuz.FullName)
		valueArgs = append(valueArgs, vuz.Logo)
		valueArgs = append(valueArgs, vuz.City)
		valueArgs = append(valueArgs, vuz.Locality)
		valueInsertCount += 5
	}
	smt := `INSERT INTO vuzes (shortname, fullname, logo, city, locality) VALUES %s`
	smt = fmt.Sprintf(smt, strings.Join(valueStrings, ","))
	tx, _ := db.Begin()
	_, err = db.Exec(smt, valueArgs...)
	if err != nil {
		tx.Rollback()
		panic(err)
	} else {
		tx.Commit()
		fmt.Println("added many vuzes...")
	}
}

func AddVuz(vuz models.Vuz) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	checkErr(err)
	defer db.Close()
	smt := `INSERT INTO vuzes (shortname, fullname, logo, city, locality) VALUES ($1, $2, $3, $4, $5)`
	tx, _ := db.Begin()
	_, err = db.Exec(smt, vuz.ShortName, vuz.FullName, vuz.Logo, vuz.City, vuz.Locality)
	if err != nil {
		tx.Rollback()
		panic(err)
	} else {
		tx.Commit()
		fmt.Println("added vuz:", vuz.FullName)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}