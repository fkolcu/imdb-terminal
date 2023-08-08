package scraping

import (
	"github.com/fkolcu/imdb-terminal/internal/util"
	"golang.org/x/net/html"
)

type ImdbPersonFoundItem struct {
	Name        string
	Poster      string
	Detail      string
	Description string
}

var personFound ImdbPersonFoundItem

func ViewImdbPerson(url string) ImdbPersonFoundItem {
	personFound = ImdbPersonFoundItem{}

	htmlTree := scrape(ImdbNameScraper{Url: url})

	itemsToResolve := []ItemResolvingAttribute{
		{
			Tag:     "span",
			Class:   "sc-afe43def-1 fDTGTb",
			OnFound: onPersonNameFound,
		},
		{
			Tag:     "div",
			Class:   "ipc-poster ipc-poster--baseAlt ipc-poster--dynamic-width ipc-sub-grid-item ipc-sub-grid-item--span-2",
			OnFound: onPersonPosterFound,
		},
		{
			Tag:     "ul",
			Class:   "ipc-inline-list ipc-inline-list--show-dividers sc-afe43def-4 kdXikI baseAlt",
			OnFound: onPersonDetailFound,
		},
		{
			Tag:     "div",
			Class:   "ipc-html-content ipc-html-content--baseAlt ipc-html-content--display-inline",
			OnFound: onPersonDescriptionFound,
		},
	}

	traverseTree(htmlTree, itemsToResolve)

	return personFound
}

func onPersonNameFound(node *html.Node, data string) {
	personFound.Name = data
}

func onPersonPosterFound(node *html.Node, data string) {
	var posterUrl string
	for _, attr := range node.FirstChild.FirstChild.Attr {
		if attr.Key == "src" {
			posterUrl = attr.Val
			break
		}
	}

	if posterUrl != "" {
		personFound.Poster = util.RetrieveImage(posterUrl)
	}
}

func onPersonDetailFound(node *html.Node, data string) {
	node = node.FirstChild

	for {
		if personFound.Detail == "" {
			personFound.Detail = node.FirstChild.Data
		} else {
			personFound.Detail += " | " + node.FirstChild.Data
		}

		if node.NextSibling == nil {
			break
		}

		node = node.NextSibling
	}
}

func onPersonDescriptionFound(node *html.Node, data string) {
	personFound.Description = node.FirstChild.FirstChild.Data
}
