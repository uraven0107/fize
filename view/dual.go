package view

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Dual struct {
	View
	left   Component
	right  Component
	keyMap map[pattern]func(Component)
	pre    tcell.Key
}

func NewDual(app *tview.Application, left Component, right Component) *Dual {
	dual := &Dual{
		left:   left,
		right:  right,
		keyMap: make(map[pattern]func(Component)),
		pre:    0,
	}
	dual.app = app
	dual.ui = tview.NewGrid().
		SetBorders(true).
		SetSize(1, 2, 0, 0)
	return dual
}

func (dual *Dual) Init() error {
	if err := dual.left.Init(); err != nil {
		return err
	}
	if err := dual.right.Init(); err != nil {
		return err
	}
	return nil
}

func (dual *Dual) Render() tview.Primitive {
	return dual.ui.(*tview.Grid).
		AddItem(dual.left.Render(), 0, 0, 1, 1, 0, 0, true).
		AddItem(dual.right.Render(), 0, 1, 1, 1, 0, 0, true)
}

func (dual *Dual) HasFocus() bool {
	return dual.ui.HasFocus()
}

func (dual *Dual) SetFocus() {
	dual.app.SetFocus(dual.ui)
}

var FocusToRight = func(c Component) {
	dual := c.(*Dual)
	if dual.left.HasFocus() {
		dual.right.SetFocus()
	}
}

var FocusToLeft = func(c Component) {
	dual := c.(*Dual)
	if dual.right.HasFocus() {
		dual.left.SetFocus()
	}
}

func (dual *Dual) MappingKey(prefix tcell.Key, key rune, fn func(Component)) {
	pattern := pattern{
		prefix: prefix,
		key:    key,
	}
	dual.keyMap[pattern] = fn
}

func (dual *Dual) InitKeyBind() {
	dual.ui.(*tview.Grid).SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if dual.pre == 0 {
			if event.Key() == tcell.KeyRune {
				pattern := pattern{
					prefix: 0,
					key:    event.Rune(),
				}
				if fn, ok := dual.keyMap[pattern]; ok {
					fn(dual)
					dual.pre = 0
					return nil
				}
				return event
			} else {
				dual.pre = event.Key()
				return nil
			}
		} else {
			if event.Key() == tcell.KeyRune {
				pattern := pattern{
					prefix: dual.pre,
					key:    event.Rune(),
				}
				if fn, ok := dual.keyMap[pattern]; ok {
					fn(dual)
					dual.pre = 0
					return nil
				}
				dual.pre = 0
				return nil
			} else {
				dual.pre = event.Key()
				return nil
			}
		}
	})
}
