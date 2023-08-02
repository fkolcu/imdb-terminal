package view

import "github.com/rivo/tview"

const (
	ButtonOk int = iota
)

type ModalConfig struct {
	Text       string
	ButtonType int
	EventOk    func()
}

func NewModal(config ModalConfig) *tview.Modal {
	myModal := tview.NewModal()
	myModal.SetText(config.Text)

	var buttons []string

	switch config.ButtonType {
	case ButtonOk:
		buttons = []string{"OK"}
	}

	myModal.AddButtons(buttons)
	myModal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonLabel == "OK" {
			config.EventOk()
		}
	})

	return myModal
}
