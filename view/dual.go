package view

import "github.com/rivo/tview"

type Dual struct {
	ui    *tview.Grid
	left  Component
	right Component
}

func NewDual(left Component, right Component) *Dual {
	return &Dual{
		left:  left,
		right: right,
	}
}

func (dual *Dual) Init() error {
	if err := dual.right.Init(); err != nil {
		return err
	}
	if err := dual.left.Init(); err != nil {
		return err
	}
	dual.ui = dual.GetLayout().(*tview.Grid)
	return nil
}

func (dual *Dual) GetLayout() tview.Primitive {
	return tview.NewGrid().
		SetBorders(true).
		SetSize(1, 2, 0, 0)
}

func (dual *Dual) Render() tview.Primitive {
	return dual.ui.
		AddItem(dual.left.Render(), 0, 0, 1, 1, 0, 0, true).
		AddItem(dual.right.Render(), 0, 1, 1, 1, 0, 0, true)
}
