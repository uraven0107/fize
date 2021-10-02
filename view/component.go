package view

import (
	"github.com/rivo/tview"
)

type Component interface {
	Init() error
	InitLayout()
	Render() tview.Primitive
}

type Bindable interface {
	MappingKey(key rune, fn func(Component))
	InitKeyBind()
}

type View struct {
	ui tview.Primitive
}
