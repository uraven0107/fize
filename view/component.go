package view

import (
	"github.com/rivo/tview"
)

type Component interface {
	GetLayout() tview.Primitive
	Render() tview.Primitive
	Init() error
}

type Bindable interface {
	MappingKey(key rune, fn func(Component))
	InitKeyBind()
}
