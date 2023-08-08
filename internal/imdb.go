package internal

import (
	"github.com/fkolcu/imdb-terminal/internal/panel"
	"github.com/fkolcu/imdb-terminal/internal/scraping"
	"github.com/fkolcu/imdb-terminal/internal/util"
	"github.com/fkolcu/imdb-terminal/internal/view"
	"github.com/rivo/tview"
	"strings"
)

var global *AppGlobal

var mainPanel panel.Panel
var mainPage *tview.Grid

var titleUrls []string
var peopleUrls []string

type CliPages struct {
	MainPage  *tview.Grid
	MainPanel panel.Panel

	TitleDetailPage  *tview.Grid
	TitleDetailPanel panel.Panel

	PersonDetailPage  *tview.Grid
	PersonDetailPanel panel.Panel
}

func NewCli(appGlobal *AppGlobal) CliPages {
	global = appGlobal
	mainPanel = panel.MainPanel{
		OnInputDone:      onSearch,
		OnTitleSelected:  onTitleSelected,
		OnPersonSelected: onPersonSelected,
	}
	mainPage = panel.NewPanel(mainPanel.GetAttributes(), mainPanel.InitializeTabs())

	appGlobal.ActivePanel = mainPanel

	return CliPages{
		MainPage:  mainPage,
		MainPanel: mainPanel,
	}
}

func newTitlePage(titleDetailPanel panel.TitleDetailPanel) CliPages {
	titleDetailPage := panel.NewPanel(titleDetailPanel.GetAttributes(), titleDetailPanel.InitializeTabs())

	return CliPages{
		TitleDetailPage:  titleDetailPage,
		TitleDetailPanel: titleDetailPanel,
	}
}

func newPersonPage(personDetailPanel panel.PersonDetailPanel) CliPages {
	personDetailPage := panel.NewPanel(personDetailPanel.GetAttributes(), personDetailPanel.InitializeTabs())

	return CliPages{
		PersonDetailPage:  personDetailPage,
		PersonDetailPanel: personDetailPanel,
	}
}

func onSearch(searchInput *tview.InputField, titleList *tview.List, personList *tview.List) {
	textInput := strings.TrimSpace(searchInput.GetText())
	if textInput == "" {
		modal := view.NewModal(view.ModalConfig{
			Text:       "You need to enter something",
			ButtonType: view.ButtonOk,
			EventOk: func() {
				global.CliUiApp.SetRoot(mainPage, true)
			},
		})
		global.CliUiApp.SetRoot(modal, false)
		return
	}

	titles, people := scraping.SearchImdb(textInput)

	titleUrls = []string{}
	peopleUrls = []string{}

	// Since range doesn't guarantee the order, for loop is best here
	for index := 0; index < len(titles); index++ {
		titleUrls = append(titleUrls, titles[index].Url)
		titleList.AddItem(titles[index].Name, titles[index].Summary, rune('0'+index), nil)
	}

	for index := 0; index < len(people); index++ {
		peopleUrls = append(peopleUrls, people[index].Url)
		personList.AddItem(people[index].Name, people[index].Summary, rune('0'+index), nil)
	}
}

func onTitleSelected(selectedIndex int) {
	titleUrl := titleUrls[selectedIndex]
	foundTitle := scraping.ViewImdbTitle(titleUrl)

	titleDetailPanel := panel.TitleDetailPanel{
		TitleName:         foundTitle.Name,
		TitlePoster:       foundTitle.Poster,
		TitleDetail:       foundTitle.Detail,
		TitleGenre:        foundTitle.Genre,
		TitleDescription:  foundTitle.Description,
		TitleMetadata:     foundTitle.Metadata,
		TitleRating:       foundTitle.Rating,
		TitleUrl:          titleUrl,
		OnSeeImdbSelected: onImdbButtonSelected,
		OnGoBackSelected:  onGoBackButtonSelected,
	}

	p := newTitlePage(titleDetailPanel)

	global.ActivePanel = titleDetailPanel
	global.CliUiApp.SetRoot(p.TitleDetailPage, true)
	ToggleFocus(global)
}

func onImdbButtonSelected(url string) {
	util.OpenUrl(url)
}

func onGoBackButtonSelected() {
	FocusOn(mainPanel, global)
	global.ActivePanel = mainPanel
	global.CliUiApp.SetRoot(mainPage, true)
}

func onPersonSelected(selectedIndex int) {
	personUrl := peopleUrls[selectedIndex]
	foundPerson := scraping.ViewImdbPerson(personUrl)

	personDetailPanel := panel.PersonDetailPanel{
		PersonName:        foundPerson.Name,
		PersonPoster:      foundPerson.Poster,
		PersonDetail:      foundPerson.Detail,
		PersonDescription: foundPerson.Description,
		PersonUrl:         personUrl,
		OnSeeImdbSelected: onImdbButtonSelected,
		OnGoBackSelected:  onGoBackButtonSelected,
	}

	p := newPersonPage(personDetailPanel)

	global.ActivePanel = personDetailPanel
	global.CliUiApp.SetRoot(p.PersonDetailPage, true)
	ToggleFocus(global)
}
