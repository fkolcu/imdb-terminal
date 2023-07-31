package main

import (
	"github.com/fkolcu/imdb-terminal/internal"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	panel := internal.DrawPanel(app)
	internal.ToggleFocus(app, panel)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEscape:
			app.Stop()
		case tcell.KeyTab:
			internal.ToggleFocus(app, panel)
		}
		return event
	})

	err := app.SetRoot(panel.Grid, true).Run()
	if err != nil {
		panic(err)
	}
}
