package main

import (
	"github.com/fkolcu/imdb-terminal/internal"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	grid := internal.DrawPanel()

	err := app.SetRoot(grid, true).Run()
	if err != nil {
		panic(err)
	}
}
