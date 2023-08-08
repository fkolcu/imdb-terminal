package panel

import (
	"github.com/fkolcu/imdb-terminal/internal/scraping"
	"github.com/fkolcu/imdb-terminal/internal/view"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strings"
)

var titleDetailTabs []PanelTab

type TitleDetailPanel struct {
	TitleName         string
	TitlePoster       string
	TitleDetail       string
	TitleGenre        string
	TitleDescription  string
	TitleMetadata     scraping.ImdbTitleMetadata
	TitleRating       scraping.ImdbTitleRating
	TitleUrl          string
	OnGoBackSelected  func()
	OnSeeImdbSelected func(url string)
}

func (t TitleDetailPanel) GetAttributes() Attributes {
	return Attributes{
		Rows:    []int{0, 0, 3, 1},
		Columns: []int{-1, -1, -1, -1},
	}
}

func (t TitleDetailPanel) InitializeTabs() []PanelTab {
	// Tab 1 : Poster
	//emptyImage := util.RetrieveEmptyImage()
	image := view.NewImage(t.TitlePoster)
	imageProperty := TabProperty{0, 0, 3, 2, 0, 0, false, false}
	imageTab := PanelTab{image, imageProperty}

	// Tab 2: Info
	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow)
	flex.SetBorderPadding(0, 1, 0, 0)

	textConfig := view.TextConfig{Border: true, BorderColor: tcell.ColorPink, Padding: view.TextPadding{Left: 1}}

	titleText := view.NewText(textConfig)
	titleText.SetText(t.TitleName)
	titleText.SetTextAlign(tview.AlignCenter)
	titleText.SetBackgroundColor(tcell.ColorFireBrick)
	flex.AddItem(titleText, 3, 1, false)

	presentationSection := tview.NewFlex()
	presentationSection.SetDirection(tview.FlexColumn)

	detailText := view.NewText(textConfig)
	detailText.SetText(getDetailText(t))
	presentationSection.AddItem(detailText, 0, 1, false)

	ratingText := view.NewText(textConfig)
	ratingText.SetText(getRatingText(t))
	presentationSection.AddItem(ratingText, 0, 1, false)
	flex.AddItem(presentationSection, 4, 1, false)

	descriptionText := view.NewText(view.TextConfig{
		Border:      true,
		BorderColor: tcell.ColorPink,
		Padding:     view.TextPadding{Left: 1},
	})
	descriptionText.SetText(t.TitleDescription)
	flex.AddItem(descriptionText, 0, 1, false)

	metadataDirectors := view.NewText(textConfig)
	metadataDirectors.SetText(strings.Join(t.TitleMetadata.Directors, ", "))
	metadataDirectors.SetTitle("Director(s)").SetTitleAlign(tview.AlignLeft)
	flex.AddItem(metadataDirectors, 3, 1, false)

	metadataWriters := view.NewText(textConfig)
	metadataWriters.SetText(strings.Join(t.TitleMetadata.Writers, ", "))
	metadataWriters.SetTitle("Writers").SetTitleAlign(tview.AlignLeft)
	flex.AddItem(metadataWriters, 3, 1, false)

	metadataStars := view.NewText(textConfig)
	metadataStars.SetText(strings.Join(t.TitleMetadata.Stars, ", "))
	metadataStars.SetTitle("Stars").SetTitleAlign(tview.AlignLeft)
	flex.AddItem(metadataStars, 3, 1, false)

	flexProperty := TabProperty{0, 2, 2, 2, 0, 0, false, false}
	flexTab := PanelTab{flex, flexProperty}

	// Tab 3 : Back Button
	backButton := view.NewButton(view.ButtonConfig{
		Label: "Go Back",
		OnSelected: func() {
			// descriptionText.SetText("Go Back Selected")
			t.OnGoBackSelected()
		},
	})
	backButtonProperty := TabProperty{2, 2, 1, 1, 0, 0, false, true}
	backButtonTab := PanelTab{backButton, backButtonProperty}

	// Tab 4 : Link Button
	linkButton := view.NewButton(view.ButtonConfig{
		Label: "See on IMDb",
		OnSelected: func() {
			// descriptionText.SetText("See On IMDb Selected")
			t.OnSeeImdbSelected(t.TitleUrl)
		},
	})
	linkButtonProperty := TabProperty{2, 3, 1, 1, 0, 0, false, true}
	linkButtonTab := PanelTab{linkButton, linkButtonProperty}

	// Tab 5: Footer
	footer := view.NewFooter()
	footerProperty := TabProperty{3, 0, 1, 4, 0, 0, false, false}
	footerTab := PanelTab{footer, footerProperty}

	titleDetailTabs = []PanelTab{imageTab, flexTab, backButtonTab, linkButtonTab, footerTab}
	return titleDetailTabs
}

func getDetailText(t TitleDetailPanel) string {
	var result string
	if t.TitleDetail != "" {
		result = t.TitleDetail + "\n"
	}

	result += t.TitleGenre

	if result != "" {
		return result
	}

	return "No details found"
}

func getRatingText(t TitleDetailPanel) string {
	if t.TitleRating.RatingValue == "" {
		return "No rating found"
	} else {
		return "⭐️ " + t.TitleRating.RatingValue + "/" + t.TitleRating.MaxValue + "\n\t" + t.TitleRating.TotalRatingNumber
	}
}

func (t TitleDetailPanel) GetTabs() []PanelTab {
	if len(titleDetailTabs) == 0 {
		t.InitializeTabs()
	}

	return titleDetailTabs
}
