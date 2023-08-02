package view

import (
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
}

func NewText(config TextConfig) *tview.TextView {
	myText := tview.NewTextView()
	myText.SetDynamicColors(config.DynamicColors)
	myText.SetBorderPadding(config.Padding.Top, config.Padding.Bottom, config.Padding.Left, config.Padding.Right)
	return myText
}
