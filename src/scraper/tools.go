package scraper

import (
	"regexp"
	"time"

	"github.com/gocolly/colly"
)

var ReTextBetweenParentheses = regexp.MustCompile(`\(([^)]*)\)`)
var ReFindDigits = regexp.MustCompile(`\d+`)

const Domain = "https://academkin.ru"

var Headers = map[string]string{
	"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
}

func getBody(url string) (body *colly.HTMLElement) {
	c := colly.NewCollector()
	c.SetRequestTimeout(30 * time.Second)
	c.OnHTML("div.page-wrap", func(h *colly.HTMLElement) {
		body = h
	})
	c.Post(url, Headers)
	return
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
