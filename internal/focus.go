package internal

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var currentTab = 0

func ToggleFocus(app *tview.Application, panel Panel) {
	switch currentTab {
	case 0:
		fallthrough
	case 3:
		app.SetFocus(panel.Tab1)
		panel.Tab1.SetBorderColor(tcell.ColorRed)
		panel.Tab2.SetBorderColor(tcell.ColorDefault)
		panel.Tab3.SetBorderColor(tcell.ColorDefault)
		currentTab = 1
	case 1:
		app.SetFocus(panel.Tab2)
		panel.Tab2.SetBorderColor(tcell.ColorRed)
		panel.Tab1.SetBorderColor(tcell.ColorDefault)
		panel.Tab3.SetBorderColor(tcell.ColorDefault)
		currentTab = 2
	case 2:
		app.SetFocus(panel.Tab3)
		panel.Tab3.SetBorderColor(tcell.ColorRed)
		panel.Tab1.SetBorderColor(tcell.ColorDefault)
		panel.Tab2.SetBorderColor(tcell.ColorDefault)
		currentTab = 3
	}
}
