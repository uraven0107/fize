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

	panel := view.NewPanel("/home/uraven")
	if err := panel.Init(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	primitive := panel.Render()
	if err := app.SetRoot(primitive, true).SetFocus(primitive).Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
