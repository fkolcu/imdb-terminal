package view

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TextPadding struct {
	Top    int
	Bottom int
	Left   int
	Right  int
}

type TextConfig struct {
	DynamicColors bool
	Padding       TextPadding
	Border        bool
	BorderColor   tcell.Color
	Alignment     int
}

func NewText(config TextConfig) *tview.TextView {
	myText := tview.NewTextView()
	myText.SetDynamicColors(config.DynamicColors)
	myText.SetBorderPadding(config.Padding.Top, config.Padding.Bottom, config.Padding.Left, config.Padding.Right)
	myText.SetWordWrap(true)

	myText.SetBorder(config.Border)
	myText.SetBorderColor(config.BorderColor)

	myText.SetTextAlign(config.Alignment)

	return myText
}
