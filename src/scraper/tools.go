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
	c.SetRequestTimeout(60 * time.Second)
	c.OnHTML("div.page-wrap", func(h *colly.HTMLElement) {
		body = h
	})
	c.Post(url, Headers)
	return
}

func setSpecsToGroups(count int, specs []string) (groups [][]string) {
	for i := 0; i < len(specs); i += count {
		group := specs[i:]
		if len(group) >= count {
			group = group[:count]
		}
		groups = append(groups, group)
	}
	return
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
