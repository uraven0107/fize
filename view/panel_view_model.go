package view

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/uraven0107/fize/view/service"
)

type PanelViewModel struct {
	panel     *Panel
	dirPath   string
	fileInfos []os.FileInfo
	keyMap    map[rune]func(*PanelViewModel)
}

func NewPanelViewModel() *PanelViewModel {
	return &PanelViewModel{
		panel:  NewPanel(),
		keyMap: make(map[rune]func(*PanelViewModel)),
	}
}

func (pvm *PanelViewModel) Init(dirPath string) error {
	fileInfos, err := service.FetchFileInfos(dirPath)
	if err != nil {
		return err
	}
	pvm.dirPath = dirPath
	pvm.fileInfos = fileInfos
	return nil
}

func (pvm *PanelViewModel) Render() tview.Primitive {
	table := pvm.panel.GetLayout().(*tview.Table)
	for i, fileInfo := range pvm.fileInfos {
		cell := tview.NewTableCell(fileInfo.Name())
		if fileInfo.IsDir() {
			cell.SetTextColor(tcell.Color120)
		}
		table.SetCell(i, 0, cell.SetReference(fileInfo))
	}
	return table
}

func (pvm *PanelViewModel) Clear() {
	table := pvm.panel.GetLayout().(*tview.Table)
	table.Clear()
}

func (pvm *PanelViewModel) Reflesh() error {
	pvm.Clear()
	err := pvm.Init(pvm.dirPath)
	if err != nil {
		return err
	}
	pvm.Render()
	return nil
}

func (pvm *PanelViewModel) MappingKey(key rune, binder func(*PanelViewModel)) {
	pvm.keyMap[key] = binder
	pvm.bindKey()
}

func (pvm *PanelViewModel) bindKey() {
	table := pvm.panel.GetLayout().(*tview.Table)
	for key, binder := range pvm.keyMap {
		table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if key == event.Rune() {
				binder(pvm)
				return nil
			}
			return event
		})
	}
}

var Reflesh = func(pvm *PanelViewModel) {
	if err := pvm.Reflesh(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var UpDir = func(pvm *PanelViewModel) {
	table := pvm.panel.GetLayout().(*tview.Table)
	row, col := table.GetSelection()
	selected := table.GetCell(row, col)
	fileInfo := selected.GetReference().(os.FileInfo)
	if fileInfo.IsDir() {
		dirPath := pvm.dirPath + "/" + fileInfo.Name()
		pvm.Init(dirPath)
		if err := pvm.Reflesh(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
