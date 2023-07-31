package internal

import (
	"golang.org/x/net/html"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ImdbItem struct {
	Name    string
	Summary string
	Url     string
}

var items = make(map[string]map[int]ImdbItem)

// Resets when item type changes
var currentIndex = 0

// Keeps the item type in the loop
var currentItemType string

// SearchImdb searches IMDb using the searchText you give it and
// returns the list of Titles and People
func SearchImdb(searchText string) (map[int]ImdbItem, map[int]ImdbItem) {
	searchText = strings.TrimSpace(searchText)
	if searchText == "" {
		return map[int]ImdbItem{}, map[int]ImdbItem{}
	}

	htmlTree := readImdb(searchText)
	traverseTree(htmlTree)

	titles := items["Titles"]
	people := items["People"]
	return titles, people
}

func readImdb(searchText string) *html.Node {
	var requestUrl = "https://www.imdb.com/find/?q=" + url.QueryEscape(searchText)

	request, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		panic(err)
	}

	header := request.Header
	header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36")

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

func traverseTree(node *html.Node) {
	resolveImdbItems(node, "h3", "ipc-title__text")
	resolveImdbItems(node, "a", "ipc-metadata-list-summary-item__t")
	resolveImdbItems(node, "span", "ipc-metadata-list-summary-item__li")

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		traverseTree(c)
	}
}

func resolveImdbItems(node *html.Node, tag string, clsAttr string) {
	if node.Type != html.ElementNode {
		return
	}

	if node.Data != tag {
		return
	}

	for _, attr := range node.Attr {
		if attr.Key == "class" && strings.Contains(attr.Val, clsAttr) {
			child := node.FirstChild
			data := child.Data
			if tag == "h3" && (data == "Titles" || data == "People") {
				currentIndex = -1
				currentItemType = data
				return
			}

			if tag == "a" {
				currentIndex++
				if len(items[currentItemType]) == 0 {
					items[currentItemType] = make(map[int]ImdbItem)
				}

				item := ImdbItem{Name: data}

				for _, attr2 := range node.Attr {
					if attr2.Key == "href" {
						item.Url = "https://www.imdb.com" + attr2.Val
					}
				}

				items[currentItemType][currentIndex] = item

				return
			}

			if tag == "span" {
				itemsToBeUpdated := items[currentItemType][currentIndex]
				if itemsToBeUpdated.Summary == "" {
					itemsToBeUpdated.Summary = data
				} else {
					itemsToBeUpdated.Summary += " | " + data
				}
				items[currentItemType][currentIndex] = itemsToBeUpdated
				return
			}
		}
	}
}
