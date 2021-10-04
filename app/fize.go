package fize

import (
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

	left := view.NewPanel("/home/uraven")
	right := view.NewPanel("/")
	dual := view.NewDual(left, right)
	root := view.NewRoot(dual)
	if err := root.Init(); err != nil {
		return err
	}
	rootView := root.Render()

	if err := app.SetRoot(rootView, true).SetFocus(rootView).Run(); err != nil {
		return err
	}

	return nil
}
