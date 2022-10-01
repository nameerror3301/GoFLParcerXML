package main

import (
	socks "GoFLParcerXML/pkg"
	"flag"
	"fmt"
	"os"
)

var (
	tgBotToken string
)

// Function for getting a token
func init() {
	flag.StringVar(&tgBotToken, "tgBotToken", "", "Telegram Bot Token")
	flag.Parse()
	if tgBotToken == "" {
		fmt.Println("-tgBotToken is not defined")
		os.Exit(1)
	}
}

func main() {
	socks.GetXmlItem("https://hidemy.name/ru/proxy-list/?type=5#list", "https://www.fl.ru/rss/all.xml")
}
