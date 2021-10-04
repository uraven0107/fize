package view

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/uraven0107/fize/utils"
	"github.com/uraven0107/fize/view/service"
)

type Panel struct {
	View
	dirPath   string
	fileInfos []os.FileInfo
	keyMap    map[rune]func(Component)
}

func NewPanel(dirPath string) Component {
	return &Panel{
		dirPath: dirPath,
		keyMap:  make(map[rune]func(Component)),
	}
}

func (panel *Panel) InitLayout() {
	panel.ui = tview.NewTable().
		SetBorders(false).
		SetSelectable(true, false).
		SetSelectedStyle(tcell.StyleDefault.Background(tcell.Color201))
}

func (panel *Panel) Render() tview.Primitive {
	for i, fileInfo := range panel.fileInfos {
		cell := tview.NewTableCell(fileInfo.Name())
		if fileInfo.IsDir() {
			cell.SetTextColor(tcell.Color120)
		}
		panel.ui.(*tview.Table).SetCell(i, 0, cell.SetReference(fileInfo))
	}
	return panel.ui
}

func (panel *Panel) Init() error {
	if err := panel.changeDir(panel.dirPath); err != nil {
		return err
	}
	panel.MappingKeyDefault()
	panel.InitKeyBind()
	return nil
}

func (panel *Panel) changeDir(dirPath string) error {
	fileInfos, err := service.FetchFileInfos(dirPath)
	if err != nil {
		return err
	}
	panel.dirPath = dirPath
	panel.fileInfos = fileInfos
	return nil
}

func (panel *Panel) clear() {
	table := panel.ui.(*tview.Table)
	table.Clear()
}

func (panel *Panel) reflesh() error {
	panel.clear()
	err := panel.changeDir(panel.dirPath)
	if err != nil {
		return err
	}
	panel.Render()
	return nil
}

func (panel *Panel) MappingKey(key rune, fn func(Component)) {
	panel.keyMap[key] = fn
}

func (panel *Panel) InitKeyBind() {
	panel.ui.(*tview.Table).SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		for key, fn := range panel.keyMap {
			if key == event.Rune() {
				fn(panel)
				return nil
			}
		}
		return event
	})
}

func (panel *Panel) MappingKeyDefault() {
	panel.MappingKey('r', Reflesh)
	panel.MappingKey('l', DownDir)
	panel.MappingKey('h', UpDir)
}

var Reflesh = func(c Component) {
	if err := c.(*Panel).reflesh(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var DownDir = func(c Component) {
	panel := c.(*Panel)
	table := panel.ui.(*tview.Table)
	row, col := table.GetSelection()
	selected := table.GetCell(row, col)
	fileInfo := selected.GetReference().(os.FileInfo)
	if fileInfo.IsDir() {
		dirPath := utils.ResolvePath(panel.dirPath, fileInfo.Name())
		panel.changeDir(dirPath)
		if err := panel.reflesh(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

var UpDir = func(c Component) {
	panel := c.(*Panel)
	dirPath := utils.ResolveParentDirPath(panel.dirPath)
	panel.changeDir(dirPath)
	if err := panel.reflesh(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
