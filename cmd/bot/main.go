package main

import (
	x "GoFLParcerXML/pkg"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
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
	// Timer 30 sec...
	t := time.NewTimer(30 * time.Second)
	fmt.Printf("Начало ожидания - %v\n", time.Now().Format(time.UnixDate))
	<-t.C
	category, title, description, link, pubDate := x.GetXmlItem("https://hidemy.name/ru/proxy-list/?type=5#list", "https://www.fl.ru/rss/all.xml")
	message := map[string]interface{}{
		"chat_id": "983044040", // Your telegram chatid
		"text":    fmt.Sprintf("Категория - [%s]\nЗаголовок - [%s]\nОписание - [\t%s]\nСсылка - [%s]\n [%s]\n", category, title, description, link, pubDate),
	}
	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalf("Err marshal message - %s\n", err)
	}
	_, err = http.Post(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?", tgBotToken), "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalf("Err sending message to telegram - %s\n", err)
	}
}
