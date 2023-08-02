package main

import (
	"github.com/fkolcu/imdb-terminal/internal"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	imdbCli := internal.NewCli(app)
	internal.ToggleFocus(app, imdbCli.MainPanel)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEscape:
			app.Stop()
		case tcell.KeyTab:
			internal.ToggleFocus(app, imdbCli.MainPanel)
		}
		return event
	})

	err := app.SetRoot(imdbCli.MainPage, true).Run()
	if err != nil {
		panic(err)
	}
}
