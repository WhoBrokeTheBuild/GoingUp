package goingup

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var app *App

// App represents the whole goingup application
type App struct {
	Options   AppOptions
	Pages     []Page
	Templates *template.Template
	Content   map[string]string
}

// AppOptions is the container for all the global application settings
type AppOptions struct {
	Port            int
	TemplateDir     string
	ContentDir      string
	StaticAssetsDir string
	StaticAssetsURL string

	LoginAction    string
	RegisterAction string

	Menus map[string][]MenuItem
}

// NewApp creates a new App instance
func NewApp() *App {
	app = &App{
		Options: AppOptions{
			Port:            80,
			TemplateDir:     "templates",
			ContentDir:      "content",
			StaticAssetsDir: "static/",
			StaticAssetsURL: "/static/",
			LoginAction:     "/login",
			RegisterAction:  "/register",
			Menus:           make(map[string][]MenuItem, 10),
		},
	}
	return app
}

// AddPage _
func (a *App) AddPage(url string, title string, tmpl string, ctnt string) error {
	if url == "" {
		return fmt.Errorf("Cannot create page with no URL")
	}

	if tmpl == "" {
		tmpl = "page"
	}

	a.Pages = append(a.Pages, Page{
		URL:      url,
		Title:    title,
		Template: tmpl,
		ContentName:  ctnt,
	})

	return nil
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
		r.HandleFunc(page.URL, makePageHandler(page))
	}

	strPort := strconv.Itoa(a.Options.Port)
	fmt.Printf("Listening on %s\n", strPort)
	http.ListenAndServe(":"+strPort, newLogHandler(r))
}
