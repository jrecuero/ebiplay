package engine

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/jrecuero/ebiplay/pkg/tools"
	"golang.org/x/image/font/basicfont"
)

type Menu struct {
	name          string
	x, y          float64
	width, height int
	menuItems     []*MenuItem
	menuLabels    []string
	menuItemIndex int
	scroller      *Scroller
	parent        *Menu
}

// NewTopMenu function creates a new Menu instance.
func NewTopMenu(name string, x, y float64, w, h int, menuItems []*MenuItem, menuItemIndex int) *Menu {
	numberOfMenuItems := len(menuItems)
	// Look for the menu item with the largest string.
	maxItemLength := 0
	for _, item := range menuItems {
		maxItemLength = tools.Max(maxItemLength, len(item.GetLabel()))
	}
	// Reassign the maximum menu item length if the horizontal size is greater
	// than the number of items by the maximum number of character for any menu
	// item.
	if (maxItemLength * numberOfMenuItems) < (w - 2) {
		maxItemLength = (w - 2) / numberOfMenuItems
	}
	// Add padding for every menu item to fill the whole horizontal length.
	paddingMenuItems := make([]string, numberOfMenuItems)
	var paddingLength float64 = float64(maxItemLength) - 1
	menuItemX := x + 1
	menuItemY := y + 1
	for i, menuItem := range menuItems {
		menuItem.SetPosition(menuItemX, menuItemY)
		paddingMenuItems[i] = fmt.Sprintf("%-*s", int(paddingLength), menuItem.GetLabel())
		menuItemX = menuItemX + paddingLength
	}
	// Assign the total number of characters required to contains all menu
	// items.
	totalSelectionLength := numberOfMenuItems * maxItemLength
	menu := &Menu{
		menuItems:     menuItems,
		menuLabels:    paddingMenuItems,
		menuItemIndex: menuItemIndex,
		parent:        nil,
	}
	for _, item := range menuItems {
		item.SetMenu(menu)
	}
	menu.scroller = NewScroller(totalSelectionLength, w-2, maxItemLength)
	return menu
}

func NewSubMenu(name string, x, y float64, w, h int, menuItems []*MenuItem, menuItemIndex int, parent *Menu) *Menu {
	selectionsLength := len(menuItems)
	// Add padding for every menu item to fill the whole horizontal length.
	paddingSelections := make([]string, selectionsLength)
	menuItemX := x + 1
	menuItemY := y + 1
	for i, menuItem := range menuItems {
		menuItem.SetPosition(menuItemX, menuItemY)
		paddingSelections[i] = fmt.Sprintf("%-*s", w-2, menuItem.GetLabel())
		menuItemY = menuItemY + 1
	}
	menu := &Menu{
		menuItems:     menuItems,
		menuLabels:    paddingSelections,
		menuItemIndex: menuItemIndex,
		parent:        parent,
		scroller:      NewVerticalScroller(selectionsLength, h-2),
	}
	return menu
}

// -----------------------------------------------------------------------------
// Menu private methods
// -----------------------------------------------------------------------------

func (m *Menu) getMenuItemLabel(index int) string {
	return m.menuLabels[index]
}

func (m *Menu) execute(args ...any) {
	switch args[0].(string) {
	case "up":
		if m.parent != nil {
			m.prevMenuItem()
		}
	case "left":
		if m.parent == nil {
			m.prevMenuItem()
		}
	case "down":
		if m.parent != nil {
			m.nextMenuItem()
		}
	case "right":
		if m.parent == nil {
			m.nextMenuItem()
		}
	case "run":
		menuItem := m.menuItems[m.menuItemIndex]
		if callback, args := menuItem.GetCallback(); callback != nil {
			if args != nil {
				args = append([]any{menuItem}, args...)
				//callback(m, args...)
			} else {
				callback(m, menuItem)
			}
		}
	}
}

func (m *Menu) nextMenuItem() {
	index := m.menuItemIndex
	for index < (len(m.menuItems) - 1) {
		index++
		if m.menuItems[index].IsEnabled() {
			m.menuItemIndex = index
			return
		}
	}
}

func (m *Menu) prevMenuItem() {
	index := m.menuItemIndex
	for index > 0 {
		index--
		if m.menuItems[index].IsEnabled() {
			m.menuItemIndex = index
			return
		}
	}
}

func (m *Menu) updateTopMenuCanvas() {
	//m.scroller.Update(m.menuItemIndex)
	//canvas := m.GetCanvas()
	//canvas.WriteRectangleInCanvasAt(nil, nil, m.GetStyle(), engine.CanvasRectSingleLine)
	//m.scroller.CreateIter()
	//for y := 1; m.scroller.IterHasNext(); {
	//    index, x := m.scroller.IterGetNext()
	//    selection := m.getMenuItemLabel(index)
	//    if index == m.menuItemIndex {
	//        reverseStyle := tools.ReverseStyle(m.GetStyle())
	//        canvas.WriteStringInCanvasAt(selection, reverseStyle, api.NewPoint(x+1, y))
	//    } else {
	//        style := m.GetStyle()
	//        if !m.menuItems[index].IsEnabled() {
	//            style = tools.SetAttrToStyle(style, tcell.AttrDim)
	//        }
	//        canvas.WriteStringInCanvasAt(selection, style, api.NewPoint(x+1, y))
	//    }
	//}
}

func (m *Menu) updateSubMenuCanvas() {
	//m.scroller.Update(m.menuItemIndex)
	//canvas := m.GetCanvas()
	//canvas.WriteRectangleInCanvasAt(nil, nil, m.GetStyle(), engine.CanvasRectSingleLine)
	//m.scroller.CreateIter()
	//for x := 1; m.scroller.IterHasNext(); {
	//    index, y := m.scroller.IterGetNext()
	//    selection := m.getMenuItemLabel(index)
	//    if index == m.menuItemIndex {
	//        canvas.WriteStringInCanvasAt(selection, m.GetStyle(), api.NewPoint(x, y+1))
	//    } else {
	//        reverseStyle := tools.ReverseStyle(m.GetStyle())
	//        if !m.menuItems[index].IsEnabled() {
	//            reverseStyle = tools.SetAttrToStyle(reverseStyle, tcell.AttrDim)
	//        }
	//        canvas.WriteStringInCanvasAt(selection, reverseStyle, api.NewPoint(x, y+1))
	//    }
	//}
}

// updateCanvas method updates the list box canvas with proper menuItems to be
// displayed and the proper selected option.
func (m *Menu) updateCanvas() {
	if m.parent == nil {
		m.updateTopMenuCanvas()
	} else {
		m.updateSubMenuCanvas()
	}
}

func (m *Menu) Draw(screen *ebiten.Image) {
	face := basicfont.Face7x13
	//for i, line := range m.menuLabels {
	//    y := 20 + i*20
	//    text.Draw(screen, line, face, 10, y, color.White)
	//}
	m.scroller.Update(m.menuItemIndex)
	m.scroller.CreateIter()
	for i := 0; m.scroller.IterHasNext(); i++ {
		j := 60 + i*10
		index, y := m.scroller.IterGetNext()
		_ = y
		selection := m.getMenuItemLabel(index)
		if index == m.menuItemIndex {
			text.Draw(screen, selection, face, 10, j, color.Black)
		} else {
			text.Draw(screen, selection, face, 10, j, color.White)
		}
	}

}

// -----------------------------------------------------------------------------
// Menu public methods
// -----------------------------------------------------------------------------

// DisableMenuItemForIndex method disables all menu items for given indexes.
func (m *Menu) DisableMenuItemForIndex(indexes ...int) error {
	for _, index := range indexes {
		if index < len(m.menuItems) {
			m.menuItems[index].SetEnabled(false)
		} else {
			return fmt.Errorf("Index %d out of range for menu %s", index, m.name)
		}
	}
	return nil
}

// DisableMenuItemForLabel method disables all menu item for given labels.
func (m *Menu) DisableMenuItemsForLabel(labels ...string) error {
	for _, label := range labels {
		if menuItem := m.FindMenuItemByLabel(label); menuItem != nil {
			menuItem.SetEnabled(false)
		} else {
			return fmt.Errorf("Label %s not found for menu %s", label, m.name)
		}
	}
	return nil
}

// EnableMenuItemForIndex method enables all menu items for given indexes.
func (m *Menu) EnableMenuItemForIndex(indexes ...int) error {
	for _, index := range indexes {
		if index < len(m.menuItems) {
			m.menuItems[index].SetEnabled(true)
		} else {
			return fmt.Errorf("Index %d out of range for menu %s", index, m.name)
		}
	}
	return nil
}

// EnableMenuItemForLabel method enables all menu item for given labels.
func (m *Menu) EnableMenuItemsForLabel(labels ...string) error {
	for _, label := range labels {
		if menuItem := m.FindMenuItemByLabel(label); menuItem != nil {
			menuItem.SetEnabled(true)
		} else {
			return fmt.Errorf("Label %s not found for menu %s", label, m.name)
		}
	}
	return nil
}

// FindMenuItemByLabel method finds the menu item for the given label.
func (m *Menu) FindMenuItemByLabel(label string) *MenuItem {
	for _, menuItem := range m.menuItems {
		if menuItem.GetLabel() == label {
			return menuItem
		}
	}
	return nil
}

// GetSelection method returns the option for the selected index.
func (m *Menu) GetSelection() string {
	return strings.TrimSpace(m.getMenuItemLabel(m.menuItemIndex))
}

func (m *Menu) Refresh() {
	m.updateCanvas()
}

// SetSelectionToIndex method sets the menu item index selected to the given
// index.
func (m *Menu) SetSelectionToIndex(index int) error {
	if index < len(m.menuItems) {
		m.menuItemIndex = index
		return nil
	}
	return fmt.Errorf("Index %d out of range for menu", index)
}

// SetSelectionToLabel method sets the menu item index selected to the given
// label.
func (m *Menu) SetSelectionToLabel(label string) error {
	for index, menuLabel := range m.menuLabels {
		if menuLabel == label {
			m.menuItemIndex = index
			return nil
		}
	}
	return fmt.Errorf("Label %s not found for menu", label)
}

// Update method executes all listbox functionality every tick time. Keyboard
// inut is scanned in order to move the selection index and proceed to select
// any option.
func (m *Menu) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		m.nextMenuItem()
	} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		m.prevMenuItem()
	}
}
