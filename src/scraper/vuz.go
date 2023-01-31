package scraper

import (
	"github.com/SalomanYu/Academkin/src/models"
	"strings"
)


func GetVuz(url string) (vuz models.Vuz){
	body := getBody(url)
	vuz.VuzId = strings.Split(url, "/")[5]
	vuz.FullName = body.ChildText("div.page-header h1 span")
	vuz.ShortName = body.ChildText("div.page-header h1 small")
	vuz.Locality = body.ChildText("td[itemprop=streetAddress]")
	vuz.City = body.ChildText("td[itemprop=addressLocality]")
	vuz.Logo = body.ChildAttr("img.high-school-avatar", "src")
	if len(vuz.Logo) > 0 {
		vuz.Logo = Domain + vuz.Logo
	}
	return
}
