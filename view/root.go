package view

import "github.com/rivo/tview"

type Root struct {
	View
	child Component
}

func NewRoot(child Component) Component {
	return &Root{
		child: child,
	}
}

func (root *Root) Init() error {
	root.InitLayout()
	if err := root.child.Init(); err != nil {
		return err
	}
	return nil
}

func (root *Root) InitLayout() {
	root.child.InitLayout()
}

func (root *Root) Render() tview.Primitive {
	return tview.NewFrame(root.child.Render())
}
