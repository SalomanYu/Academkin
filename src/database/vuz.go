package database

import (
	"fmt"
	"github.com/SalomanYu/Academkin/src/models"

	_ "github.com/lib/pq"
)

func (d *Database) SaveVuz(vuz models.Vuz) {
	query := fmt.Sprintf("INSERT INTO academkin_vuz(id, short_name, full_name, image, city, address, url) VALUES ($1, $2, $3, $4, $5, $6, $7) ON CONFLICT DO NOTHING;")
	tx, _ := d.db.Begin()
	_, err := d.db.Exec(query, vuz.VuzId, vuz.ShortName, vuz.FullName, vuz.Logo, vuz.City, vuz.Locality, vuz.Url)
	checkErr(err)
	tx.Commit()
	fmt.Printf("Успешно сохранили вуз:%s \n", vuz.FullName)
}
