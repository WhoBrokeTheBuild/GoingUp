package goingup

// AppOptions is the container for all the global application settings
type AppOptions struct {
	Port            int
	TemplateDir     string
	StaticAssetsDir string
	StaticAssetsURL string

	LoginAction    string
	RegisterAction string

	Menus map[string][]MenuItem
}
