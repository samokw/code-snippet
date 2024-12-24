package main

import (
	"html/template"
	"path/filepath"

	"code-snippet.samokw/internal/models"
)

type templateData struct {
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

		tplate, err := template.ParseFiles("./ui/html/base.html")
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
