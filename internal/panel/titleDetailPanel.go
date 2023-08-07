package panel

import (
	"github.com/fkolcu/imdb-terminal/internal/view"
	"github.com/rivo/tview"
)

var titleDetailTabs []PanelTab

type TitleDetailPanel struct {
	TitleName         string
	TitlePoster       string
	TitleDetail       string
	TitleGenre        string
	TitleDescription  string
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

	titleText := view.NewText(view.TextConfig{}).SetText(t.TitleName)
	flex.AddItem(titleText, 0, 1, false)

	detailText := view.NewText(view.TextConfig{}).SetText(t.TitleDetail)
	flex.AddItem(detailText, 0, 1, false)

	genreText := view.NewText(view.TextConfig{}).SetText(t.TitleGenre)
	flex.AddItem(genreText, 0, 1, false)

	descriptionText := view.NewText(view.TextConfig{}).SetText(t.TitleDescription)
	flex.AddItem(descriptionText, 0, 1, false)

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

func (t TitleDetailPanel) GetTabs() []PanelTab {
	if len(titleDetailTabs) == 0 {
		t.InitializeTabs()
	}

	return titleDetailTabs
}
