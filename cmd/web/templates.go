package main

import (
	"html/template"
	"path/filepath"
)

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.gohtml"))
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}
		ts, err = template.ParseGlob(filepath.Join(dir, "*partial.gohtml"))
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return cache, nil
}
