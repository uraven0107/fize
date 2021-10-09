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
	keyMap    map[pattern]func(Component)
	pre       tcell.Key
}

func NewPanel(app *tview.Application, dirPath string) *Panel {
	panel := &Panel{
		dirPath: dirPath,
		keyMap:  make(map[pattern]func(Component)),
		pre:     0,
	}
	panel.app = app
	panel.ui = tview.NewTable().
		SetBorders(false).
		SetSelectable(true, false).
		SetSelectedStyle(tcell.StyleDefault.Background(tcell.Color201))
	panel.focused = func() {
		table := panel.ui.(*tview.Table)
		table.SetBackgroundColor(tcell.Color237)
	}
	panel.unfocused = func() {
		table := panel.ui.(*tview.Table)
		table.SetBackgroundColor(0)
	}
	return panel
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

func (panel *Panel) MappingKey(prefix tcell.Key, key rune, fn func(Component)) {
	pattern := pattern{
		prefix: prefix,
		key:    key,
	}
	panel.keyMap[pattern] = fn
}

func (panel *Panel) InitKeyBind() {
	panel.ui.(*tview.Table).SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if panel.pre == 0 {
			if event.Key() == tcell.KeyRune {
				pattern := pattern{
					prefix: 0,
					key:    event.Rune(),
				}
				if fn, ok := panel.keyMap[pattern]; ok {
					fn(panel)
					panel.pre = 0
					return nil
				}
				return event
			} else {
				panel.pre = event.Key()
				return nil
			}
		} else {
			if event.Key() == tcell.KeyRune {
				pattern := pattern{
					prefix: panel.pre,
					key:    event.Rune(),
				}
				if fn, ok := panel.keyMap[pattern]; ok {
					fn(panel)
					panel.pre = 0
					return nil
				}
				panel.pre = 0
				return nil
			} else {
				panel.pre = event.Key()
				return nil
			}
		}
	})
}

var Reflesh = func(c Component) {
	panel := c.(*Panel)
	if err := panel.reflesh(); err != nil {
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
