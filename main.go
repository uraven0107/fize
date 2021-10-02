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

	left := view.NewPanel("/home/uraven")
	right := view.NewPanel("/")
	dual := view.NewDual(left, right)
	if err := dual.Init(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	primitive := dual.Render()

	if err := app.SetRoot(primitive, true).SetFocus(primitive).Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
