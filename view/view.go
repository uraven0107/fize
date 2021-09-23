package view

import "github.com/rivo/tview"

type View struct {
	viewModel *ViewModel
}

func NewView(viewModel *ViewModel) *View {
	return &View{
		viewModel: viewModel,
	}
}

func (view *View) Render() tview.Primitive {
	table := tview.NewTable().SetBorders(false)
	for i, fileInfo := range view.viewModel.fileInfos {
		table.SetCell(i, 0, tview.NewTableCell(fileInfo.Name()))
	}
	return table
}
