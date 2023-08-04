package internal

import (
	"github.com/fkolcu/imdb-terminal/internal/panel"
	"github.com/fkolcu/imdb-terminal/internal/scraping"
	"github.com/fkolcu/imdb-terminal/internal/view"
	"github.com/rivo/tview"
	"strings"
)

var app *tview.Application

var mainPanel panel.Panel
var mainPage *tview.Grid

type CliPages struct {
	MainPage  *tview.Grid
	MainPanel panel.Panel

	TitleDetailPage  *tview.Grid
	TitleDetailPanel panel.Panel
}

func NewCli(application *tview.Application) CliPages {
	app = application

	mainPanel = panel.MainPanel{
		OnInputDone: onSearch,
	}
	mainPage = panel.NewPanel(mainPanel.GetAttributes(), mainPanel.GetTabs())

	return CliPages{
		MainPage:  mainPage,
		MainPanel: mainPanel,
	}
}

func onSearch(searchInput *tview.InputField, titleList *tview.List, personList *tview.List) {
	textInput := strings.TrimSpace(searchInput.GetText())
	if textInput == "" {
		modal := view.NewModal(view.ModalConfig{
			Text:       "You need to enter something",
			ButtonType: view.ButtonOk,
			EventOk: func() {
				app.SetRoot(mainPage, true)
			},
		})
		app.SetRoot(modal, false)
		return
	}

	titles, people := scraping.SearchImdb(textInput)

	// Since range doesn't guarantee the order, for loop is best here
	for index := 0; index < len(titles); index++ {
		//titleUrls = append(titleUrls, titles[index].Url)
		titleList.AddItem(titles[index].Name, titles[index].Summary, rune('0'+index), nil)
	}

	for index := 0; index < len(people); index++ {
		//peopleUrls = append(peopleUrls, people[index].Url)
		personList.AddItem(people[index].Name, people[index].Summary, rune('0'+index), nil)
	}
}
