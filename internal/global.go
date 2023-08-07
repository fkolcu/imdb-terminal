package internal

import (
	"github.com/fkolcu/imdb-terminal/internal/panel"
	"github.com/rivo/tview"
)

type AppGlobal struct {
	CliUiApp    *tview.Application
	ActivePanel panel.Panel
}
