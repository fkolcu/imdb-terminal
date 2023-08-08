package scraping

import (
	"golang.org/x/net/html"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ImdbScraper interface {
	GetUrl() string
}

type ImdbSearchScraper struct {
	SearchText string
}

func (iss ImdbSearchScraper) GetUrl() string {
	urlSafeSearchText := url.QueryEscape(iss.SearchText)
	return "https://www.imdb.com/find/?q=" + urlSafeSearchText
}

type ImdbNameScraper struct {
	Url string
}

func (ins ImdbNameScraper) GetUrl() string {
	return ins.Url
}

type ImdbTitleScraper struct {
	Url string
}

func (its ImdbTitleScraper) GetUrl() string {
	return its.Url
}

const (
	UserAgent = "Mozilla/5.0 (Windows' NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36"
)

func scrape(scraper ImdbScraper) *html.Node {
	request, err := http.NewRequest("GET", scraper.GetUrl(), nil)
	if err != nil {
		panic(err)
	}

	header := request.Header
	header.Set("User-Agent", UserAgent)

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	htmlTree, err := html.Parse(response.Body)
	if err != nil {
		panic(err)
	}

	return htmlTree
}

type ItemResolvingAttribute struct {
	Tag     string
	Class   string
	OnFound func(node *html.Node, data string)
}

func traverseTree(node *html.Node, itemsToResolve []ItemResolvingAttribute) {
	for i := 0; i < len(itemsToResolve); i++ {
		resolveItems(node, itemsToResolve[i])
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		traverseTree(c, itemsToResolve)
	}
}

func resolveItems(node *html.Node, itemToResolve ItemResolvingAttribute) {
	if node.Type != html.ElementNode {
		return
	}

	if node.Data != itemToResolve.Tag {
		return
	}

	if hasClass(node, itemToResolve.Class) {
		child := node.FirstChild

		var data string
		if child != nil {
			data = child.Data
		} else {
			data = node.Data
		}

		itemToResolve.OnFound(node, data)
	}
}

type ScrapeNode struct {
	*html.Node
}

func hasClass(node *html.Node, className string) bool {
	for _, attribute := range node.Attr {
		if attribute.Key == "class" {
			if strings.Contains(attribute.Val, className) {
				return true
			}
			return false
		}
	}

	return false
}
