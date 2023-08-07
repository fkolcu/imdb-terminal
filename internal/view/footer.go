package view

import "github.com/rivo/tview"

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

	// more to come later

	return myFooter
}
