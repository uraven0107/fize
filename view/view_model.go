package view

import (
	"github.com/rivo/tview"
)

type ViewModel interface {
	Render() tview.Primitive
}
