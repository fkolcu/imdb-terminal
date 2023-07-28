package internal

import (
	"fmt"
	"github.com/rivo/tview"
)

func buildSearch() *tview.InputField {
	return NewInput(InputConfig{
		Placeholder: "Search IMDb",
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
	fmt.Fprintf(textView, "%s", "[yellow]Press Tab to switch between tabs")
	return textView
}

func DrawPanel() *tview.Grid {
	grid := tview.NewGrid()

	// There are 3 rows and 2 columns
	// First one will have 1 item
	// Second one (at the middle) will have 2 items
	// Last one will have 1 item
	grid.SetRows(2, 0, 1)
	grid.SetColumns(-1, -1)
	grid.SetBorders(true)

	// Row 1 Column 1-2
	grid.AddItem(buildSearch(), 0, 0, 1, 2, 0, 0, true)

	// Row 2 Column 1
	grid.AddItem(buildTitles([]ListItem{}), 1, 0, 1, 1, 0, 0, false)

	// Row 2 Column 2
	grid.AddItem(buildPeople([]ListItem{}), 1, 1, 1, 1, 0, 0, false)

	// Row 3 Column 1-2
	grid.AddItem(buildFooter(), 2, 0, 1, 2, 0, 0, false)

	//searchBox := tview.NewFlex()
	//searchBox.SetBorder(true)
	//searchBox.SetBorderColor(tcell.ColorRed)
	//searchBox.SetDirection(tview.FlexColumn)
	//searchBox.AddItem(searchInput, 0, 1, false)
	//
	//listBox := tview.NewFlex()

	return grid
}
