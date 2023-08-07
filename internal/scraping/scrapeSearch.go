package scraping

import (
	"fmt"
	"golang.org/x/net/html"
)

var currentIndex = 0
var currentItemType string

var items = make(map[string]map[int]ImdbFoundItem)

type ImdbFoundItem struct {
	Name    string
	Summary string
	Url     string
}

func SearchImdb(searchText string) (map[int]ImdbFoundItem, map[int]ImdbFoundItem) {
	htmlTree := scrape(ImdbSearchScraper{SearchText: searchText})

	itemsToResolve := []ItemResolvingAttribute{
		{
			Tag:     "h3",
			Class:   "ipc-title__text",
			OnFound: onTitleOrNameFound,
		},
		{
			Tag:     "a",
			Class:   "ipc-metadata-list-summary-item__t",
			OnFound: onItemFound,
		},
		{
			Tag:     "span",
			Class:   "ipc-metadata-list-summary-item__li",
			OnFound: onItemSummaryFound,
		},
	}

	traverseTree(htmlTree, itemsToResolve)

	if len(items) > 0 {
		fmt.Println("")
	}

	return items["Titles"], items["People"]
}

func onTitleOrNameFound(node *html.Node, data string) {
	if data != "Titles" && data != "People" {
		return
	}

	currentIndex = -1
	currentItemType = data
}

func onItemFound(node *html.Node, data string) {
	currentIndex++
	if len(items[currentItemType]) == 0 {
		items[currentItemType] = make(map[int]ImdbFoundItem)
	}

	item := ImdbFoundItem{Name: data}
	for _, attr := range node.Attr {
		if attr.Key == "href" {
			item.Url = "https://www.imdb.com" + attr.Val
			break
		}
	}

	items[currentItemType][currentIndex] = item
}

func onItemSummaryFound(node *html.Node, data string) {
	itemsToBeUpdated := items[currentItemType][currentIndex]

	if itemsToBeUpdated.Summary != "" {
		data = itemsToBeUpdated.Summary + " | " + data
	}
	itemsToBeUpdated.Summary = data

	items[currentItemType][currentIndex] = itemsToBeUpdated
}
