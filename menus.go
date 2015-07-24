package goingup

// MenuItem is a representation of a menu item
type MenuItem struct {
	URL    string
	Text   string
	Active bool
}

func setMenuActive(menu []MenuItem, url string) {
	for ind, mi := range menu {
		if url == mi.URL {
			menu[ind].Active = true
		} else {
			menu[ind].Active = false
		}
	}
}
