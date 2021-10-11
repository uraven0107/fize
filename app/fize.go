package fize

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/uraven0107/fize/view"
)

func Run() error {
	app := tview.NewApplication()

	left := view.NewPanel(app, "/home/uraven")
	left.MappingKey(0, 'r', view.Reflesh)
	left.MappingKey(0, 'l', view.DownDir)
	left.MappingKey(0, 'h', view.UpDir)

	right := view.NewPanel(app, "/")
	right.MappingKey(0, 'r', view.Reflesh)
	right.MappingKey(0, 'l', view.DownDir)
	right.MappingKey(0, 'h', view.UpDir)

	dual := view.NewDual(app, left, right)
	dual.MappingKey(tcell.KeyCtrlW, 'l', view.FocusToRight)
	dual.MappingKey(tcell.KeyCtrlW, 'h', view.FocusToLeft)
	dual.SetFocusToLeft()

	root := view.NewRoot(app, dual)

	if err := root.Init(); err != nil {
		return err
	}
	rootView := root.Render()

	if err := app.SetRoot(rootView, true).Run(); err != nil {
		return err
	}

	return nil
}
