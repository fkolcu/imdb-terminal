package internal

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ListConfig struct {
	Title               string
	ItemSelectedHandler func(int, string, string, rune)
}

type ListItem struct {
	MainText      string
	SecondaryText string
}

func NewList(items []ListItem, config ListConfig) *tview.List {
	myList := tview.NewList()

	//myList.SetTitle(config.Title)
	myList.SetBorder(true)
	myList.SetBorderColor(tcell.ColorDefault)
	myList.SetSelectedFocusOnly(true)

	for index, item := range items {
		myList.AddItem(item.MainText, item.SecondaryText, rune('0'+index), nil)
	}

	myList.SetSelectedFunc(config.ItemSelectedHandler)

	return myList
}
