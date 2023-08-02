package panel

import (
	"github.com/fkolcu/imdb-terminal/internal/view"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type MainPanel struct {
	OnInputDone func(searchInput *tview.InputField, titleList *tview.List, personList *tview.List)
}

var tabs []PanelTab

func (m MainPanel) GetAttributes() Attributes {
	return Attributes{
		Rows:    []int{3, 0, 1},
		Columns: []int{-1, -1},
	}
}

func (m MainPanel) GetTabs() []PanelTab {
	if len(tabs) == 0 {
		tabs = m.getTabs()
	}

	return tabs
}

var titleList *tview.List
var personList *tview.List

func (m MainPanel) getTabs() []PanelTab {
	// Tab 1: Search Input
	searchInput := view.NewInput(view.InputConfig{
		Placeholder: "Search IMDb (press <enter> to search)",
		Border:      true,
	})
	searchInput.SetDoneFunc(func(key tcell.Key) {
		titleList.Clear()
		personList.Clear()

		m.OnInputDone(searchInput, titleList, personList)
	})
	searchInputProperty := TabProperty{0, 0, 1, 2, 0, 0, true}
	searchInputTab := PanelTab{searchInput, searchInputProperty}

	// Tab 2: Title List
	titleList = view.NewList([]view.ListItem{}, view.ListConfig{
		Title: "Titles",
	})
	titleListProperty := TabProperty{1, 0, 1, 1, 0, 0, false}
	titleListTab := PanelTab{titleList, titleListProperty}

	// Tab 3: Person list
	personList = view.NewList([]view.ListItem{}, view.ListConfig{
		Title: "People",
	})
	personListProperty := TabProperty{1, 1, 1, 1, 0, 0, false}
	personListTab := PanelTab{personList, personListProperty}

	// Tab 4: Footer
	keybindingText := view.NewText(view.TextConfig{DynamicColors: true, Padding: view.TextPadding{Left: 1, Right: 1}})
	keybindingText.SetText("[blue]<esc>: Quit, <tab>: Jump between panels")
	keybindingTextProperty := TabProperty{2, 0, 1, 2, 0, 0, false}
	keybindingTextTab := PanelTab{keybindingText, keybindingTextProperty}

	return []PanelTab{searchInputTab, titleListTab, personListTab, keybindingTextTab}
}
