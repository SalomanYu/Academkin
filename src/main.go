package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/SalomanYu/Academkin/src/database"
	"github.com/SalomanYu/Academkin/src/logger"
	"github.com/SalomanYu/Academkin/src/scraper"
	"github.com/joho/godotenv"
)

var db *database.Database

func init() {
	logger.Log.Println("Program started.")
	err := godotenv.Load("../.env")
	if err != nil {
		panic("Создайте файл с переменными окружения! .env")
	}
}

func main() {
	s := time.Now().Unix()
	port, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	db = &database.Database{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     port,
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASS"),
		Name:     os.Getenv("POSTGRES_DBNAME"),
	}
	db.InitDatabase()
	start()

	fmt.Println(time.Now().Unix()-s, "seconds")
	logger.Log.Println(time.Now().Unix()-s, "seconds")
}

func start() {
	var pageNum int
	for {
		pageNum++
		logger.Log.Println("Main Page:", pageNum)

		vuzesUrls := scraper.GetAllVuzUrlsInPage(pageNum)
		for _, item := range vuzesUrls {
			saveVuzInfo(item)
		}
		if len(vuzesUrls) == 0 {
			logger.Log.Println("Вузы закончились на странице: ", pageNum)

			break
		}
	}
}

func saveVuzInfo(vuzUrl string) {
	vuz := scraper.GetVuz(vuzUrl)
	vuzSpecs := scraper.GetAllSpecualizations(vuzUrl)

	db.SaveVuz(vuz)
	db.SaveVuzSpecializations(vuzSpecs)
	logger.Log.Println("Save vuz:", vuz.FullName)

}
