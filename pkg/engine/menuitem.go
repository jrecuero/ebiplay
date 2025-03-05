package engine

// MenuItem structure defines every item in a menu widget.
type MenuItem struct {
	label        string
	enabled      bool
	x, y         float64
	menu         *Menu
	callback     func(...any)
	callbackArgs []any
}

func NewMenuItem(label string) *MenuItem {
	return &MenuItem{
		label:        label,
		enabled:      true,
		x:            0,
		y:            0,
		menu:         nil,
		callback:     nil,
		callbackArgs: nil,
	}
}

func NewExtendedMenuItem(label string, enabled bool, menu *Menu, callback func(...any), args []any) *MenuItem {
	return &MenuItem{
		label:        label,
		enabled:      enabled,
		x:            0,
		y:            0,
		menu:         menu,
		callback:     callback,
		callbackArgs: args,
	}
}

// -----------------------------------------------------------------------------
// MenuItem public methods
// -----------------------------------------------------------------------------

func (m *MenuItem) GetCallback() (func(...any), []any) {
	return m.callback, m.callbackArgs
}

func (m *MenuItem) GetLabel() string {
	return m.label
}

func (m *MenuItem) GetMenu() *Menu {
	return m.menu
}

func (m *MenuItem) GetPosition() (float64, float64) {
	return m.x, m.y
}

func (m *MenuItem) IsEnabled() bool {
	return m.enabled
}

func (m *MenuItem) SetCallback(calback func(...any), args []any) *MenuItem {
	m.callback = calback
	m.callbackArgs = args
	return m
}

func (m *MenuItem) SetEnabled(enabled bool) *MenuItem {
	m.enabled = enabled
	return m
}

func (m *MenuItem) SetLabel(label string) *MenuItem {
	m.label = label
	return m
}

func (m *MenuItem) SetMenu(menu *Menu) *MenuItem {
	m.menu = menu
	return m
}

func (m *MenuItem) SetPosition(x, y float64) *MenuItem {
	m.x = x
	m.y = y
	return m
}
