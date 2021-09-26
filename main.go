package main

import (
	"fmt"
	"os"

	"github.com/rivo/tview"
	"github.com/uraven0107/fize/view"
)

// FIXME 仮実装
func main() {
	app := tview.NewApplication()

	viewModel := view.NewPanelViewModel()
	viewModel.InitKeyBind()
	err := viewModel.InitDir("/home/uraven")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	primitive := viewModel.Render()
	if err := app.SetRoot(primitive, true).SetFocus(primitive).Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
