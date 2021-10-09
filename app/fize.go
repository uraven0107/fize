package fize

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/uraven0107/fize/view"
)

type component_type int

const (
	dual component_type = iota
	panel
)

func Run() error {
	app := tview.NewApplication()

	left := view.NewPanel(app, "/home/uraven")
	left.MappingKey(0, 'r', view.Reflesh)
	left.MappingKey(0, 'l', view.DownDir)
	left.MappingKey(0, 'h', view.UpDir)
	left.InitKeyBind()

	right := view.NewPanel(app, "/")
	right.MappingKey(0, 'r', view.Reflesh)
	right.MappingKey(0, 'l', view.DownDir)
	right.MappingKey(0, 'h', view.UpDir)
	right.InitKeyBind()

	dual := view.NewDual(app, left, right)
	dual.MappingKey(tcell.KeyCtrlW, 'l', view.FocusToRight)
	dual.MappingKey(tcell.KeyCtrlW, 'h', view.FocusToLeft)
	dual.InitKeyBind()

	root := view.NewRoot(app, dual)

	if err := root.Init(); err != nil {
		return err
	}
	rootView := root.Render()

	left.SetFocus()
	if err := app.SetRoot(rootView, true).Run(); err != nil {
		return err
	}

	return nil
}
