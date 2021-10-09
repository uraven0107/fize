package view

import "github.com/rivo/tview"

type Root struct {
	View
	child Component
}

func NewRoot(app *tview.Application, child Component) *Root {
	root := &Root{
		child: child,
	}
	root.app = app
	root.ui = nil // rootコンポーネントにレイアウトは存在しない
	return root
}

func (root *Root) Init() error {
	if err := root.child.Init(); err != nil {
		return err
	}
	return nil
}

func (root *Root) Render() tview.Primitive {
	return tview.NewFrame(root.child.Render())
}
