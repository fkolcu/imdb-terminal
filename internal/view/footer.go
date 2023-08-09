package view

import (
	"github.com/fkolcu/imdb-terminal/internal/util"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	bindingText = "[blue]<esc>: Quit, <tab>: Jump between tabs/buttons, <enter>: Select"
)

func NewFooter() *tview.Flex {
	myFooter := tview.NewFlex()
	myFooter.SetDirection(tview.FlexColumn)
	myFooter.SetBorderPadding(0, 0, 1, 1)

	keybindingText := NewText(TextConfig{DynamicColors: true})
	keybindingText.SetText(bindingText)
	myFooter.AddItem(keybindingText, 0, 1, false)

	githubLinkText := NewText(TextConfig{})
	githubLinkText.SetText("imdb-terminal")
	githubLinkText.SetBackgroundColor(tcell.ColorYellow)
	githubLinkText.SetTextColor(tcell.ColorBlack)
	githubLinkText.SetFocusFunc(func() {
		util.OpenUrl("https://github.com/fkolcu/imdb-terminal")
	})
	myFooter.AddItem(githubLinkText, 13, 1, false)

	// more to come later

	return myFooter
}
