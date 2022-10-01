package pkg

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func SocksParce(url string) []string {
	var arr []string
	c := colly.NewCollector()
	c.OnHTML("div > table tr", func(e *colly.HTMLElement) {
		ipText := e.DOM.Find("td:nth-child(1)").Text()
		portText := e.DOM.Find("td:nth-child(2)").Text()
		// Условие для пропуска первого значения
		if !strings.Contains(ipText, "IP адрес") && !strings.Contains(portText, "Порт") {
			arr = append(arr, fmt.Sprintf("%s:%s", ipText, portText))
		}
	})
	return arr
}
