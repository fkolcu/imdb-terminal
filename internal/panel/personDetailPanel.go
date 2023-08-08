package panel

import (
	"github.com/fkolcu/imdb-terminal/internal/view"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var personDetailTabs []PanelTab

type PersonDetailPanel struct {
	PersonName        string
	PersonPoster      string
	PersonDetail      string
	PersonDescription string
	PersonUrl         string
	OnGoBackSelected  func()
	OnSeeImdbSelected func(url string)
}

func (p PersonDetailPanel) GetAttributes() Attributes {
	return Attributes{
		Rows:    []int{0, 0, 3, 1},
		Columns: []int{-1, -1, -1, -1},
	}
}

func (p PersonDetailPanel) InitializeTabs() []PanelTab {
	// Tab 1 : Poster
	//emptyImage := util.RetrieveEmptyImage()
	image := view.NewImage(p.PersonPoster)
	imageProperty := TabProperty{0, 0, 3, 2, 0, 0, false, false}
	imageTab := PanelTab{image, imageProperty}

	// Tab 2: Info
	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow)
	flex.SetBorderPadding(0, 1, 0, 0)

	textConfig := view.TextConfig{Border: true, BorderColor: tcell.ColorPink, Padding: view.TextPadding{Left: 1}}

	personText := view.NewText(textConfig)
	personText.SetText(p.PersonName)
	personText.SetTextAlign(tview.AlignCenter)
	personText.SetBackgroundColor(tcell.ColorFireBrick)
	flex.AddItem(personText, 3, 1, false)

	detailText := view.NewText(view.TextConfig{
		Border:      true,
		BorderColor: tcell.ColorPink,
		Padding:     view.TextPadding{Left: 1},
	})
	detailText.SetText(p.PersonDetail)
	flex.AddItem(detailText, 3, 1, false)

	descriptionText := view.NewText(view.TextConfig{
		Border:      true,
		BorderColor: tcell.ColorPink,
		Padding:     view.TextPadding{Left: 1},
	})
	descriptionText.SetText(p.PersonDescription + "\n\n" + p.PersonDescription + "\n\n" + p.PersonDescription)
	descriptionText.SetScrollable(true)
	flex.AddItem(descriptionText, 0, 1, false)

	flexProperty := TabProperty{0, 2, 2, 2, 0, 0, false, false}
	flexTab := PanelTab{flex, flexProperty}

	// Tab 3 : Back Button
	backButton := view.NewButton(view.ButtonConfig{
		Label: "Go Back",
		OnSelected: func() {
			// descriptionText.SetText("Go Back Selected")
			p.OnGoBackSelected()
		},
	})
	backButtonProperty := TabProperty{2, 2, 1, 1, 0, 0, false, true}
	backButtonTab := PanelTab{backButton, backButtonProperty}

	// Tab 4 : Link Button
	linkButton := view.NewButton(view.ButtonConfig{
		Label: "See on IMDb",
		OnSelected: func() {
			// descriptionText.SetText("See On IMDb Selected")
			p.OnSeeImdbSelected(p.PersonUrl)
		},
	})
	linkButtonProperty := TabProperty{2, 3, 1, 1, 0, 0, false, true}
	linkButtonTab := PanelTab{linkButton, linkButtonProperty}

	// Tab 5: Footer
	footer := view.NewFooter()
	footerProperty := TabProperty{3, 0, 1, 4, 0, 0, false, false}
	footerTab := PanelTab{footer, footerProperty}

	personDetailTabs = []PanelTab{imageTab, flexTab, backButtonTab, linkButtonTab, footerTab}
	return personDetailTabs
}

func (p PersonDetailPanel) GetTabs() []PanelTab {
	if len(personDetailTabs) == 0 {
		p.InitializeTabs()
	}

	return personDetailTabs
}
