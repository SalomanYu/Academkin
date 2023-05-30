package scraper

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

const SearchDomain = "https://academkin.ru/vuz/list"



func GetAllVuzUrlsInPage(page int) (urls []string) {
	body := getBody(fmt.Sprintf("%s/%d", SearchDomain, page))
	body.ForEach("div.high-schools-list a.media", func(i int, h *colly.HTMLElement) {
		url := Domain + h.Attr("href")
		urls = append(urls, url)
	})
	return
}

func GetAllVuzSpecializations(vuzUrl string) (specializationUrls []string) {
	url := strings.ReplaceAll(vuzUrl, "view/", "")
	var pageNum int
	for {
		pageNum++
		findedUrls := getAllSpecializationUrlsInPage(fmt.Sprintf("%s/specialities/list/%d", url, pageNum))
		if len(findedUrls) == 0 {
			break
		}
		specializationUrls = append(specializationUrls, findedUrls...)
	}
	return
}

func getAllSpecializationUrlsInPage(url string) (urls []string) {
	body := getBody(url)
	body.ForEach("a.list-group-item", func(i int, h *colly.HTMLElement) {
		url := Domain + h.Attr("href")
		urls = append(urls, url)
	})
	return
}