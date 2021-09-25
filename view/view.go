package view

import "github.com/rivo/tview"

type View interface {
	GetLayout() tview.Primitive
}
