package main

import (
	"html/template"
	"path/filepath"
	"time"

	"code-snippet.samokw/internal/models"
)

type templateData struct {
	CurrentYear int
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		tplate, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.html")
		if err != nil {
			return nil, err
		}
		tplate, err = tplate.ParseGlob("./ui/html/components/*.html")
		if err != nil {
			return nil, err
		}
		tplate, err = tplate.ParseFiles(page)
		if err != nil {
			return nil, err
		}
		cache[name] = tplate
	}
	return cache, nil
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}