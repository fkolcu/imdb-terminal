package view

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ButtonConfig struct {
	Label      string
	OnSelected func()
}

func NewButton(config ButtonConfig) *tview.Button {
	myButton := tview.NewButton(config.Label)
	myButton.SetSelectedFunc(config.OnSelected)

	myButton.SetBorder(true)
	myButton.SetStyle(
		tcell.StyleDefault.Background(tcell.ColorGrey).
			Foreground(tview.Styles.PrimaryTextColor),
	)

	myButton.SetBackgroundColorActivated(tcell.ColorDeepPink)
	myButton.SetLabelColorActivated(tcell.ColorWhite)

	return myButton
}
