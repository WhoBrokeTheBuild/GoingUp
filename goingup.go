package goingup

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// App represents the whole goingup application
type App struct {
	Options   *Config
	Pages     []*Page
	Menus     map[string]*Menu
	Templates *template.Template
	Content   map[string]string
}

// NewApp creates a new App instance
func NewApp() *App {
	return &App{
		Options: NewConfig(),
		Pages:   make([]*Page, 0),
		Menus:   make(map[string]*Menu, 0),
	}
}

func (a *App) AddPage(page *Page) {
    a.Pages = append(a.Pages, page)
}

func (a *App) AddMenu(name string, menu *Menu) {
    a.Menus[name] = menu
}

// Run finalizes all options and calls the ListenAndServe function to serve
// requests
func (a *App) Run() {
	fmt.Println("GoingUp App Starting")

	r := mux.NewRouter()

	fs := http.FileServer(http.Dir(a.Options.StaticAssetsDir))
	r.PathPrefix(a.Options.StaticAssetsURL).Handler(http.StripPrefix(a.Options.StaticAssetsURL, fs))

	fmt.Println("Parsing Templates")
	a.Templates = template.Must(template.ParseGlob(a.Options.TemplateDir + "/*"))

	fmt.Println("Parsing Content")
	a.Content = parseContentGlob(a.Options.ContentDir + "/*.md")

	for _, page := range a.Pages {
		content := ""
		if page.ContentName != "" {
			if val, exists := a.Content[page.ContentName]; exists {
				content = val
			}
		}
		page.Content = template.HTML(content)
		r.HandleFunc(page.URL, makePageHandler(a, page))
	}

	strPort := strconv.Itoa(a.Options.Port)
	fmt.Printf("Listening on %s\n", strPort)
	http.ListenAndServe(":"+strPort, newLogHandler(r))
}

func newLogHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}