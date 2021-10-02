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
	if err := left.Init(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	right := view.NewPanel("/")
	if err := right.Init(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dual := view.NewDual(left, right)
	primitive := dual.Render()

	if err := app.SetRoot(primitive, true).SetFocus(primitive).Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
