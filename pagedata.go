package goingup

import (
    "net/http"
)

// pageData is the data that will be passed to the templates
type pageData struct {
	*Page
	*Config
    Menus map[string]*Menu
}

func makePageHandler(app *App, page *Page) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		app.Templates.ExecuteTemplate(
            rw,
            page.Template,
            pageData{
                page,
                app.Options,
                app.Menus,
            },
        )
	}
}
