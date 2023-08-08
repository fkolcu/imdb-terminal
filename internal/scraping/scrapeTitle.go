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
	Metadata    ImdbTitleMetadata
}

type ImdbTitleMetadata struct {
	Directors []string
	Writers   []string
	Stars     []string
}

var titleFound ImdbTitleFoundItem

var hasMetadataBeenFound = false

func ViewImdbTitle(url string) ImdbTitleFoundItem {
	titleFound = ImdbTitleFoundItem{}

	// This needs to be reset everytime search is done
	// since 2 metadata `ul` html tags are in IMDb, one for desktop view and other for mobile view
	hasMetadataBeenFound = false

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
		{
			Tag:     "ul",
			Class:   "ipc-metadata-list ipc-metadata-list--dividers-all title-pc-list ipc-metadata-list--baseAlt",
			OnFound: onMetaDataFound,
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

var currentMetadataType string

func onMetaDataFound(node *html.Node, data string) {
	if hasMetadataBeenFound {
		return
	}

	hasMetadataBeenFound = true

	traverseTree(node, []ItemResolvingAttribute{
		{
			Tag:     "span",
			Class:   "ipc-metadata-list-item__label ipc-metadata-list-item__label",
			OnFound: onMetaDataTypeFound,
		},
		{
			Tag:     "a",
			Class:   "ipc-metadata-list-item__label ipc-metadata-list-item__label",
			OnFound: onMetaDataTypeFound,
		},
		{
			Tag:     "a",
			Class:   "ipc-metadata-list-item__list-content-item ipc-metadata-list-item__list-content-item--link",
			OnFound: onMetaDataItemFound,
		},
	})
}

// onMetaDataTypeFound is part of onMetaDataFound
func onMetaDataTypeFound(node *html.Node, data string) {
	currentMetadataType = node.FirstChild.Data
}

// onMetaDataItemFound is part of onMetaDataFound
func onMetaDataItemFound(node *html.Node, data string) {
	switch currentMetadataType {
	case "Director", "Directors":
		titleFound.Metadata.Directors = append(titleFound.Metadata.Directors, node.FirstChild.Data)
	case "Writers":
		titleFound.Metadata.Writers = append(titleFound.Metadata.Writers, node.FirstChild.Data)
	case "Stars":
		titleFound.Metadata.Stars = append(titleFound.Metadata.Stars, node.FirstChild.Data)
	}
}
