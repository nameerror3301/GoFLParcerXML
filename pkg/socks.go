package pkg

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

// Parsing addresses from the site
func SocksParce(url string) string {
	var arr []string
	c := colly.NewCollector()
	c.OnHTML("div > table tr", func(e *colly.HTMLElement) {
		ipText := e.DOM.Find("td:nth-child(1)").Text()
		portText := e.DOM.Find("td:nth-child(2)").Text()
		if !strings.Contains(ipText, "IP адрес") && !strings.Contains(portText, "Порт") {
			arr = append(arr, fmt.Sprintf("%s:%s\n", ipText, portText))
		}
	})
	c.Visit(url)
	return arr[0]
}
