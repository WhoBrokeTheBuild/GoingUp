package goingup

import (
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

func NewPage(url string, title string, tmpl string, ctnt string) *Page {
	if url == "" {
		return nil
	}

	if tmpl == "" {
		tmpl = "page"
	}
	
	return &Page{
		URL:         url,
		Title:       title,
		Template:    tmpl,
		ContentName: ctnt,
	}
}
