package panel

import (
	"github.com/fkolcu/imdb-terminal/internal/view"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var mainTabs []PanelTab
var titleList *tview.List
var personList *tview.List

type MainPanel struct {
	OnInputDone      func(searchInput *tview.InputField, titleList *tview.List, personList *tview.List)
	OnTitleSelected  func(selectedIndex int)
	OnPersonSelected func(selectedIndex int)
}

func (m MainPanel) GetAttributes() Attributes {
	return Attributes{
		Rows:    []int{3, 0, 1},
		Columns: []int{-1, -1},
	}
}
func (m MainPanel) InitializeTabs() []PanelTab {
	// Tab 1: Search Input
	searchInput := view.NewInput(view.InputConfig{
		Placeholder: "Search IMDb (press <enter> to search)",
		Border:      true,
	})
	searchInput.SetDoneFunc(func(key tcell.Key) {
		if key != tcell.KeyEnter {
			return
		}

		titleList.Clear()
		personList.Clear()

		m.OnInputDone(searchInput, titleList, personList)
	})
	searchInputProperty := TabProperty{0, 0, 1, 2, 0, 0, true, true}
	searchInputTab := PanelTab{searchInput, searchInputProperty}

	// Tab 2: Title List
	titleList = view.NewList([]view.ListItem{}, view.ListConfig{
		Title: "Titles",
	})
	titleList.SetSelectedFunc(func(i int, s string, s2 string, r rune) {
		m.OnTitleSelected(i)
	})
	titleListProperty := TabProperty{1, 0, 1, 1, 0, 0, false, true}
	titleListTab := PanelTab{titleList, titleListProperty}

	// Tab 3: Person list
	personList = view.NewList([]view.ListItem{}, view.ListConfig{
		Title: "People",
	})
	personList.SetSelectedFunc(func(i int, s string, s2 string, r rune) {
		m.OnPersonSelected(i)
	})
	personListProperty := TabProperty{1, 1, 1, 1, 0, 0, false, true}
	personListTab := PanelTab{personList, personListProperty}

	// Tab 4: Footer
	footer := view.NewFooter()
	footerProperty := TabProperty{2, 0, 1, 2, 0, 0, false, false}
	footerTab := PanelTab{footer, footerProperty}

	mainTabs = []PanelTab{searchInputTab, titleListTab, personListTab, footerTab}
	return mainTabs
}

func (m MainPanel) GetTabs() []PanelTab {
	if len(mainTabs) == 0 {
		m.InitializeTabs()
	}

	return mainTabs
}
