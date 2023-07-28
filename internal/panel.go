package internal

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Tab interface {
	tview.Primitive
	SetBorderColor(color tcell.Color) *tview.Box
}

type Panel struct {
	Grid *tview.Grid
	Tab1 Tab
	Tab2 Tab
	Tab3 Tab
}

func buildSearch() *tview.InputField {
	return NewInput(InputConfig{
		Placeholder: "Search IMDb",
		Border:      true,
	})
}

func buildTitles(titles []ListItem) *tview.List {
	return NewList(titles, ListConfig{
		Title: "Titles",
	})
}

func buildPeople(people []ListItem) *tview.List {
	return NewList(people, ListConfig{
		Title: "People",
	})
}

func buildFooter() *tview.TextView {
	textView := tview.NewTextView()
	textView.SetDynamicColors(true)
	textView.SetBorder(true)
	textView.SetBorderColor(tcell.ColorDefault)
	fmt.Fprintf(textView, "%s", "[yellow]Press Tab to switch between tabs")
	return textView
}

func DrawPanel() Panel {
	grid := tview.NewGrid()

	// There are 3 rows and 2 columns
	// First one will have 1 item
	// Second one (at the middle) will have 2 items
	// Last one will have 1 item
	grid.SetRows(3, 0, 3)
	grid.SetColumns(-1, -1)

	// Row 1 Column 1-2
	tab1 := buildSearch()
	grid.AddItem(tab1, 0, 0, 1, 2, 0, 0, true)

	// Row 2 Column 1
	tab2 := buildTitles([]ListItem{})
	grid.AddItem(tab2, 1, 0, 1, 1, 0, 0, false)

	// Row 2 Column 2
	tab3 := buildPeople([]ListItem{})
	grid.AddItem(tab3, 1, 1, 1, 1, 0, 0, false)

	// Row 3 Column 1-2
	grid.AddItem(buildFooter(), 2, 0, 1, 2, 0, 0, false)

	return Panel{
		Grid: grid,
		Tab1: tab1,
		Tab2: tab2,
		Tab3: tab3,
	}
}
