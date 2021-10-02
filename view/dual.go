package view

import "github.com/rivo/tview"

type Dual struct {
	View
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
	dual.InitLayout()
	if err := dual.right.Init(); err != nil {
		return err
	}
	if err := dual.left.Init(); err != nil {
		return err
	}
	return nil
}

func (dual *Dual) InitLayout() {
	dual.ui = tview.NewGrid().
		SetBorders(true).
		SetSize(1, 2, 0, 0)
}

func (dual *Dual) Render() tview.Primitive {
	return dual.ui.(*tview.Grid).
		AddItem(dual.left.Render(), 0, 0, 1, 1, 0, 0, true).
		AddItem(dual.right.Render(), 0, 1, 1, 1, 0, 0, true)
}
