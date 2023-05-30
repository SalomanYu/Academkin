package database

import (
	"fmt"
	"github.com/SalomanYu/Academkin/src/models"
	"strings"
)

func (d *Database) SaveVuzSpecializations(specs []models.Specialization) {
	query, args := createQueryForMultipleInsertSpecializations(specs)
	tx, _ := d.db.Begin()
	_, err := d.db.Exec(query, args...)
	checkErr(err)
	tx.Commit()
	fmt.Printf("Успешно сохранили %d специализаций\n", len(specs))
}

func createQueryForMultipleInsertSpecializations(specs []models.Specialization) (query string, valArgs []interface{}) {
	valStrings := []string{}
	valInsertCount := 1
	for _, s := range specs {
		valStrings = append(valStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d)", valInsertCount, valInsertCount+1, valInsertCount+2, valInsertCount+3, valInsertCount+4, valInsertCount+5))
		valArgs = append(valArgs, s.Name, s.FormEducation, s.Duration, s.PreparationLevel, s.Qualification, s.Url)
		valInsertCount += 6
	}
	query = `INSERT INTO academkin_specialization(name, form, duration, level, profession, url) VALUES` + strings.Join(valStrings, ",")
	return
}
