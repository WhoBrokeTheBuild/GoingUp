package goingup

import (
	"fmt"
	"net/http"
	"html/template"
)

// Page represents a web page
type Page struct {
	Title       string
	URL         string
	Template    string
	ContentName string
	Content     template.HTML
}

// pageData is the data that will be passed to the templates
type pageData struct {
	Page
	Opts AppOptions
}

func newPageData(req *http.Request, page Page) *pageData {
	return &pageData{page, app.Options}
}

func makePageHandler(page Page) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		app.Templates.ExecuteTemplate(rw, page.Template, newPageData(req, page))
	}
}

func newLogHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
