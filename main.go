package main

import (
	"fmt"
	"os"

	"github.com/rivo/tview"
	"github.com/uraven0107/fize/view"
)

// FIXME 仮実装
func main() {
	viewModel, err := view.NewViewModel(".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	view := view.NewView(viewModel)
	if err := tview.NewApplication().SetRoot(view.Render(), true).Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
