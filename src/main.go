package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/SalomanYu/Academkin/src/database"
	"github.com/SalomanYu/Academkin/src/logger"
	"github.com/SalomanYu/Academkin/src/scraper"
)

var db database.Database

func main() {
	s := time.Now().Unix()
	db = database.Database{}
	logger.Log.Println("Program started.")
	// start()
	fmt.Println(time.Now().Unix()-s, "seconds")
	logger.Log.Println(time.Now().Unix()-s, "seconds")
	fmt.Scan(&s)
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
	specializationsUrls := scraper.GetAllVuzSpecializations(vuzUrl)
	logger.Log.Println("Count url specs:", len(specializationsUrls))
	fmt.Println("Count url specs:", len(specializationsUrls))
	var wg sync.WaitGroup
	wg.Add(len(specializationsUrls))
	for _, item := range specializationsUrls {
		go scraper.SaveSpecialization(item, &wg)
	}
	wg.Wait()

	// mongo.AddVuz(&vuz)
	logger.Log.Println("Save vuz:", vuz.FullName)

}
