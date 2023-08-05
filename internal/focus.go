package internal

import (
	"fmt"
	"github.com/fkolcu/imdb-terminal/internal/panel"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func ToggleFocus(app *tview.Application, panel panel.Panel) {
	tabs := panel.GetTabs()

	for i := 0; i < len(tabs); i++ {
		if tabs[i].Tab.HasFocus() {
			tabs[i].Tab.SetBorderColor(tcell.ColorDefault)

			focusableTab, err := findFocusableTab(tabs, i+1)
			if err == nil {
				focusableTab.Tab.SetBorderColor(tcell.ColorRed)
				app.SetFocus(focusableTab.Tab)
				return
			}

			break
		}
	}

	app.SetFocus(tabs[0].Tab)
	tabs[0].Tab.SetBorderColor(tcell.ColorRed)
}

func findFocusableTab(tabs []panel.PanelTab, startIndex int) (panel.PanelTab, error) {
	maxIndex := len(tabs) - 1
	if startIndex > maxIndex {
		startIndex = 0
	}

	for i := startIndex; i < len(tabs); i++ {
		if tabs[i].Property.Focusable == true {
			return tabs[i], nil
		}
	}

	return panel.PanelTab{}, fmt.Errorf("%s", "No focusable tab found")
}
