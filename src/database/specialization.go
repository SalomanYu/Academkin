package database

import (
	"fmt"
	"github.com/SalomanYu/Academkin/src/models"
	"strings"
)

func (d *Database) SaveVuzSpecializations(specs []models.Specialization) {
	if len(specs) == 0 {
		return
	}
	query, args := createQueryForMultipleInsertSpecializations(specs)
	tx, _ := d.db.Begin()
	_, err := d.db.Exec(query, args...)
	if err != nil {
		fmt.Println(query)
	}
	checkErr(err)
	tx.Commit()
	fmt.Printf("Успешно сохранили %d специализаций\n", len(specs))
}

func createQueryForMultipleInsertSpecializations(specs []models.Specialization) (query string, valArgs []interface{}) {
	valStrings := []string{}
	valInsertCount := 1
	for _, s := range specs {
		valStrings = append(valStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)", valInsertCount, valInsertCount+1, valInsertCount+2, valInsertCount+3, valInsertCount+4, valInsertCount+5, valInsertCount+6, valInsertCount+7))
		valArgs = append(valArgs, s.Id, s.VuzId, replaceAllQuotes(s.Name), s.FormEducation, s.Duration, strings.TrimSpace(s.PreparationLevel), s.Qualification, s.Url)
		valInsertCount += 8
	}
	query = `INSERT INTO academkin_specialization(id, vuz_id, name, form, duration, level, profession, url) VALUES` + strings.Join(valStrings, ",") + "ON CONFLICT DO NOTHING;"
	return
}
