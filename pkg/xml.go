package pkg

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/proxy"
)

// Structure for the XML file
type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Channel struct {
		Text          string `xml:",chardata"`
		Title         string `xml:"title"`
		Link          string `xml:"link"`
		Description   string `xml:"description"`
		Language      string `xml:"language"`
		PubDate       string `xml:"pubDate"`
		LastBuildDate string `xml:"lastBuildDate"`
		Docs          string `xml:"docs"`
		Generator     string `xml:"generator"`
		Image         struct {
			Text   string `xml:",chardata"`
			URL    string `xml:"url"`
			Title  string `xml:"title"`
			Link   string `xml:"link"`
			Width  string `xml:"width"`
			Height string `xml:"height"`
		} `xml:"image"`
		ManagingEditor string `xml:"managingEditor"`
		WebMaster      string `xml:"webMaster"`
		Item           []struct {
			Text        string `xml:",chardata"`
			Title       string `xml:"title"`
			Link        string `xml:"link"`
			Description string `xml:"description"`
			Guid        string `xml:"guid"`
			Category    string `xml:"category"`
			PubDate     string `xml:"pubDate"`
		} `xml:"item"`
	} `xml:"channel"`
}

// Pull the data we need
func GetXmlItem(urlSocks, urlFl string) {
	file, err := ioutil.ReadAll(Request(urlSocks, urlFl))
	if err != nil {
		log.Fatalf("Err read responce body - %s\n", err)
	}

	var r Rss
	err = xml.Unmarshal(file, &r)
	if err != nil {
		log.Fatalf("Err unmarshal xml to struct - %s\n", err)
	}
	fmt.Println(r.Channel.Item[0].Category)
	fmt.Println(r.Channel.Item[0].Title)
	fmt.Println(r.Channel.Item[0].Description)
	fmt.Println(r.Channel.Item[0].Link)
	fmt.Println(r.Channel.Item[0].PubDate)
}

// Getting XML in the body of the request
func Request(urlSocks, urlFl string) io.ReadCloser {
	ip := strings.TrimSpace(SocksParce(urlSocks))
	dialer, err := proxy.SOCKS5("tcp", ip, nil, proxy.Direct)
	if err != nil {
		log.Fatalf("Err add proxy - %s\n", err)
	}
	client := &http.Client{
		Transport: &http.Transport{
			Dial: dialer.Dial,
		},
	}
	resp, err := client.Get(urlFl)
	if err != nil {
		log.Fatalf("Err responce xml - %s\n", err)
	}
	return resp.Body
}
