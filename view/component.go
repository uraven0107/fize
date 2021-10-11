package view

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Component interface {
	Init() error
	GetLayout() tview.Primitive
	Render() tview.Primitive
	HasFocus() bool
	SetFocus()
	UnFocus()
}

type View struct {
	ui        tview.Primitive
	app       *tview.Application
	focused   func()
	unfocused func()
}

func (view *View) GetLayout() tview.Primitive {
	return view.ui
}

func (view *View) HasFocus() bool {
	return view.ui.HasFocus()
}

func (view *View) SetFocus() {
	view.app.SetFocus(view.ui)
	view.focused()
}

func (view *View) UnFocus() {
	view.ui.Blur()
	view.unfocused()
}

type pattern struct {
	prefix tcell.Key
	key    rune
}

type Bindable interface {
	MappingKey(prefix tcell.Key, key rune, fn func(Component))
	InitKeyBind()
}
