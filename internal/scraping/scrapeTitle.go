package scraping

import (
	"github.com/fkolcu/imdb-terminal/internal/util"
	"golang.org/x/net/html"
)

type ImdbTitleFoundItem struct {
	Name        string
	Poster      string
	Detail      string
	Genre       string
	Description string
}

var titleFound ImdbTitleFoundItem

func ViewImdbTitle(url string) ImdbTitleFoundItem {
	titleFound = ImdbTitleFoundItem{}

	htmlTree := scrape(ImdbTitleScraper{Url: url})

	itemsToResolve := []ItemResolvingAttribute{
		{
			Tag:     "span",
			Class:   "sc-afe43def-1 fDTGTb",
			OnFound: onTitleHeaderFound,
		},
		{
			Tag:     "div",
			Class:   "ipc-media ipc-media--poster-27x40 ipc-image-media-ratio--poster-27x40 ipc-media--baseAlt ipc-media--poster-l ipc-poster__poster-image ipc-media__img",
			OnFound: onTitlePosterFound,
		},
		{
			Tag:     "ul",
			Class:   "ipc-inline-list ipc-inline-list--show-dividers sc-afe43def-4 kdXikI baseAlt",
			OnFound: onTitleDetailFound,
		},
		{
			Tag:     "a",
			Class:   "ipc-chip ipc-chip--on-baseAlt",
			OnFound: onTitleGenreFound,
		},
		{
			Tag:     "span",
			Class:   "sc-466bb6c-2 eVLpWt",
			OnFound: onTitleDescriptionFound,
		},
	}

	traverseTree(htmlTree, itemsToResolve)

	return titleFound
}

func onTitleHeaderFound(node *html.Node, data string) {
	titleFound.Name = data
}

func onTitlePosterFound(node *html.Node, data string) {
	var posterUrl string
	for _, attr := range node.FirstChild.Attr {
		if attr.Key == "src" {
			posterUrl = attr.Val
			break
		}
	}

	if posterUrl != "" {
		titleFound.Poster = util.RetrieveImage(posterUrl)
	}
}

var itemsToResolveForDetail []ItemResolvingAttribute

func onTitleDetailFound(node *html.Node, data string) {

	if node.Data == "li" {
		if node.FirstChild.Data == "a" {
			itemsToResolveForDetail = []ItemResolvingAttribute{
				{
					Tag:     "a",
					Class:   "ipc-link ipc-link--baseAlt ipc-link--inherit-color",
					OnFound: onTitleDetailFound,
				},
			}
			traverseTree(node.FirstChild, itemsToResolveForDetail)
		} else {
			if titleFound.Detail == "" {
				titleFound.Detail = node.FirstChild.Data
			} else {
				titleFound.Detail += " - " + node.FirstChild.Data
			}
		}

		node = node.NextSibling
		if node == nil {
			return
		}
	}

	if node.Data == "a" {
		if titleFound.Detail != "" {
			data = titleFound.Detail + " - " + data
		}
		titleFound.Detail = data
		return
	}

	itemsToResolveForDetail = []ItemResolvingAttribute{
		{
			Tag:     "li",
			Class:   "ipc-inline-list__item",
			OnFound: onTitleDetailFound,
		},
	}

	if node.Data == "ul" {
		traverseTree(node, itemsToResolveForDetail)
	}
}

func onTitleGenreFound(node *html.Node, data string) {
	data = node.FirstChild.FirstChild.Data
	if titleFound.Genre != "" {
		data = titleFound.Genre + " | " + data
	}

	titleFound.Genre = data
}

func onTitleDescriptionFound(node *html.Node, data string) {
	titleFound.Description = data
}
