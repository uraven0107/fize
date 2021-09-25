package view

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Panel struct {
	layout tview.Primitive
}

func NewPanel() *Panel {
	layout := tview.NewTable().SetBorders(false).SetSelectable(true, false).SetSelectedStyle(tcell.StyleDefault.Background(tcell.Color200))
	return &Panel{
		layout: layout,
	}
}

func (panel *Panel) GetLayout() tview.Primitive {
	return panel.layout
}
