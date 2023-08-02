package internal

import (
	"github.com/fkolcu/imdb-terminal/internal/panel"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func ToggleFocus(app *tview.Application, panel panel.Panel) {
	tabs := panel.GetTabs()
	lastTabIndex := len(tabs) - 1

	var focused = false
	for i := 0; i < len(tabs); i++ {
		if tabs[i].Tab.HasFocus() {
			if i == lastTabIndex {
				app.SetFocus(tabs[0].Tab)
				tabs[0].Tab.SetBorderColor(tcell.ColorRed)
			} else {
				app.SetFocus(tabs[i+1].Tab)
				tabs[i+1].Tab.SetBorderColor(tcell.ColorRed)
			}

			tabs[i].Tab.SetBorderColor(tcell.ColorDefault)

			focused = true
			break
		}
	}

	if focused == false {
		app.SetFocus(tabs[0].Tab)
		tabs[0].Tab.SetBorderColor(tcell.ColorRed)
	}
}
