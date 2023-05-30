package scraper

import (
	"github.com/SalomanYu/Academkin/src/logger"
	"github.com/SalomanYu/Academkin/src/models"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly"
)

type TableValues struct {
	FormEducation    string
	PreparationLevel string
	Duration         string
	Qualification    string
}

func SaveSpecialization(url string, wg *sync.WaitGroup) {
	specialization := GetSpecialization(url)
	if len(specialization.Name) != 0 {
		mongo.AddSpecialization(&specialization)
		logger.Log.Println("Save spec:", specialization.Name)

	}
	wg.Done()
}

func GetSpecialization(url string) (specialization models.Specialization) {
	body := getBody(url)
	if body == nil {
		time.Sleep(1 * time.Second)
		body = getBody(url)
		if body == nil {
			logger.Log.Println("Не смогли спарсить специализацию:", url)
			return
		}
	}
	header := body.ChildText("div.page-header h1")
	specialization.Id = getSpecializationId(header)
	specialization.VuzId = strings.Split(url, "/")[4]
	specialization.Name = strings.ReplaceAll(header, ReTextBetweenParentheses.FindString(header), "")
	specialization.VuzFullName = prepareVuzTitleForSpecialization(body.ChildText("div.page-header h4"))

	tableValues := getTableValues(body)
	specialization.Duration = tableValues.Duration
	specialization.Qualification = tableValues.Qualification
	specialization.FormEducation = tableValues.FormEducation
	specialization.Url = url
	specialization.PreparationLevel = tableValues.PreparationLevel
	return
}

func prepareVuzTitleForSpecialization(title string) (result string) {
	without_br := strings.ReplaceAll(title, "\n", "")
	without_space := strings.ReplaceAll(without_br, "\t", "")
	if strings.Contains(without_space, ",") {
		result = strings.Split(without_space, ",")[1]
	} else {
		result = without_space
	}
	return strings.TrimSpace(result)
}

func getSpecializationId(text string) (id int) {
	id, _ = strconv.Atoi(ReFindDigits.FindString(text))
	return
}

func getTableValues(body *colly.HTMLElement) (values TableValues) {
	body.ForEach("table.table.table-striped tr", func(i int, row *colly.HTMLElement) {
		var title string
		row.ForEach("td", func(i int, h *colly.HTMLElement) {
			if i == 0 {
				title = h.Text
			} else if i == 1 {
				switch title {
				case "Форма обучения:":
					values.FormEducation = h.Text
				case "Срок обучения:":
					values.Duration = h.Text
				case "Уровень подготовки:":
					values.PreparationLevel = h.Text
				case "Квалификация:":
					values.Qualification = h.Text
				}
			}
		})
	})
	return
}
