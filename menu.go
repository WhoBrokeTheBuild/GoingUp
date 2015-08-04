package goingup

type Menu struct {
	Items []*MenuItem
}

func NewMenu() *Menu {
	return &Menu{ }
}

func (m *Menu) AddItem(item *MenuItem) {
	m.Items = append(m.Items, item)
}

// MenuItem is a representation of a menu item
type MenuItem struct {
	URL    string
	Text   string
}

func NewMenuItem(url string, text string) *MenuItem {
	return &MenuItem{
		URL: url,
		Text: text,
	}
}