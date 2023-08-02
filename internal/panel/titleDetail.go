package panel

import (
	"github.com/fkolcu/imdb-terminal/internal/util"
	"github.com/fkolcu/imdb-terminal/internal/view"
	"github.com/rivo/tview"
)

type TitleDetailPanel struct {
}

func (t TitleDetailPanel) GetAttributes() Attributes {
	return Attributes{
		Rows:    []int{0, 0},
		Columns: []int{-1, -1},
	}
}

func (t TitleDetailPanel) GetTabs() []PanelTab {
	if len(tabs) == 0 {
		tabs = t.getTabs()
	}

	return tabs
}

func (t TitleDetailPanel) getTabs() []PanelTab {
	// Tab 1 : Poster
	emptyImage := util.RetrieveEmptyImage()
	image := view.NewImage(emptyImage)
	imageProperty := TabProperty{0, 0, 1, 1, 0, 0, false}
	imageTab := PanelTab{image, imageProperty}

	// Tab 2: Info
	flex := tview.NewFlex()

	titleText := view.NewText(view.TextConfig{}).SetText("Title")
	flex.AddItem(titleText, 0, 1, false)

	detailText := view.NewText(view.TextConfig{}).SetText("2023 - PG-13 - 1h 55m")
	flex.AddItem(detailText, 0, 1, false)

	genreText := view.NewText(view.TextConfig{}).SetText("Genre1 | Genre2")
	flex.AddItem(genreText, 0, 1, false)

	descriptionText := view.NewText(view.TextConfig{}).SetText("This is the description of the title.")
	flex.AddItem(descriptionText, 0, 1, false)

	flexProperty := TabProperty{0, 0, 1, 1, 0, 0, false}
	flexTab := PanelTab{flex, flexProperty}

	return []PanelTab{imageTab, flexTab}
}
