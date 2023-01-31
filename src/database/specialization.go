// Пока этот пакет не используется, потому что нужно разобраться со всеми ошибками, которые возникают при записи в SQL

package database

import (
	"github.com/SalomanYu/Academkin/src/models"
	"fmt"
	"database/sql"
	"strings"
)

func AddMultipleSpecializations(specs []models.Specialization) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	checkErr(err)
	defer db.Close()
	valueStrings := []string{}
	valueArgs := []interface{}{}
	valueInsertCount := 1
	for _, spec := range specs {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d)", valueInsertCount, valueInsertCount+1, valueInsertCount+2, valueInsertCount+3, valueInsertCount+4, valueInsertCount+5, valueInsertCount+6))
		valueArgs = append(valueArgs, spec.VuzFullName)
		valueArgs = append(valueArgs, spec.Id)
		valueArgs = append(valueArgs, spec.Name)
		valueArgs = append(valueArgs, spec.FormEducation)
		valueArgs = append(valueArgs, spec.Duration)
		valueArgs = append(valueArgs, spec.PreparationLevel)
		valueArgs = append(valueArgs, spec.Qualification)
		valueInsertCount += 7
	}
	smt := `INSERT INTO specializations (vuz_fullname, id, fullname, form_education, duration, preparation_level, qualification) VALUES %s`
	smt = fmt.Sprintf(smt, strings.Join(valueStrings, ","))
	tx, _ := db.Begin()
	_, err = db.Exec(smt, valueArgs...)
	if err != nil {
		tx.Rollback()
		panic(err)
	} else {
		tx.Commit()
		fmt.Printf("Added %d specializations\n", len(specs))
	}
}

func AddSpecialization(spec models.Specialization) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	checkErr(err)
	defer db.Close()
	smt := `INSERT INTO specializations (vuz_fullname, id, fullname, form_education, duration, preparation_level, qualification) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	tx, _ := db.Begin()
	_, err = db.Exec(smt, spec.VuzFullName, spec.Id, spec.Name, spec.FormEducation, spec.Duration, spec.PreparationLevel, spec.Qualification)
	if err != nil {
 		// panic(err)
		fmt.Println("Не записался:", spec.Name)
		return
	} else {
		tx.Commit()
		fmt.Println("added spec:", spec.Name)
	}
}