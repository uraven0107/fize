package view

import "github.com/rivo/tview"

type Dual struct {
	layout *tview.Grid
	left   Component
	right  Component
}

func NewDual(left Component, right Component) *Dual {
	return &Dual{
		layout: tview.NewGrid().SetBorders(true).SetSize(1, 2, 0, 0),
		left:   left,
		right:  right,
	}
}

func (dual *Dual) GetLayout() tview.Primitive {
	return dual.layout
}

func (dual *Dual) Render() tview.Primitive {
	return dual.layout.
		AddItem(dual.left.Render(), 0, 0, 1, 1, 0, 0, true).
		AddItem(dual.right.Render(), 0, 1, 1, 1, 0, 0, true)
}
