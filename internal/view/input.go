package view

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type InputConfig struct {
	Placeholder string
	Border      bool
}

func NewInput(config InputConfig) *tview.InputField {
	myInput := tview.NewInputField()

	myInput.SetFieldBackgroundColor(tcell.ColorWhite)
	myInput.SetFieldTextColor(tcell.ColorBlack)

	myInput.SetBorder(config.Border)
	myInput.SetBorderColor(tcell.ColorDefault)

	myInput.SetPlaceholder(config.Placeholder)
	myInput.SetPlaceholderStyle(
		myInput.GetPlaceholderStyle().
			Background(tcell.ColorWhite).
			Foreground(tcell.ColorGrey),
	)

	return myInput
}
