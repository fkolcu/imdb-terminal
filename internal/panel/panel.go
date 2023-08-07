package panel

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Panel interface {
	GetAttributes() Attributes
	InitializeTabs() []PanelTab
	GetTabs() []PanelTab
}

type Tab interface {
	tview.Primitive
	SetBorderColor(color tcell.Color) *tview.Box
}

type PanelTab struct {
	Tab      Tab
	Property TabProperty
}

type TabProperty struct {
	RowSize        int
	ColumnSize     int
	RowSpan        int
	ColSpan        int
	MinPanelHeight int
	MinPanelWidth  int
	Focus          bool
	Focusable      bool
}

type Attributes struct {
	Rows    []int
	Columns []int
}

func NewPanel(attributes Attributes, tabs []PanelTab) *tview.Grid {
	myPanel := tview.NewGrid()
	myPanel.SetRows(attributes.Rows...)
	myPanel.SetColumns(attributes.Columns...)

	for i := 0; i < len(tabs); i++ {
		tab := tabs[i]
		myPanel.AddItem(
			tab.Tab,
			tab.Property.RowSize,
			tab.Property.ColumnSize,
			tab.Property.RowSpan,
			tab.Property.ColSpan,
			tab.Property.MinPanelHeight,
			tab.Property.MinPanelWidth,
			tab.Property.Focus,
		)
	}

	return myPanel
}
