package main

import (
	"github.com/fkolcu/imdb-terminal/internal"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	app.EnableMouse(true)

	appGlobal := &internal.AppGlobal{
		CliUiApp: app,
	}

	imdbCli := internal.NewCli(appGlobal)
	internal.ToggleFocus(appGlobal)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEscape:
			app.Stop()
		case tcell.KeyTab:
			internal.ToggleFocus(appGlobal)
		}
		return event
	})

	err := app.SetRoot(imdbCli.MainPage, true).Run()
	if err != nil {
		panic(err)
	}
}
