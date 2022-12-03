package ast

import (
	"github.com/kapitanov/notion2html/internal/html"
)

type DBEntry struct {
	ID         string
	Properties []*DBEntryProperty
	Content    *DBEntryContent
	Title      string
}

func (ast *DBEntry) NewProperty(name string, value DBEntryPropertyValue) *DBEntryProperty {
	entry := &DBEntryProperty{
		Name:  name,
		Value: value,
	}
	ast.Properties = append(ast.Properties, entry)
	return entry
}

func (ast *DBEntry) Render(w *html.Writer) {
	ast.Title = "Details"
	for _, prop := range ast.Properties {
		t, ok := prop.Value.(*DBEntryPropertyTitleValue)
		if ok {
			ast.Title = t.Value
		}
	}

	if ast.Content != nil {
		ast.Content.Render(w)
	}

	for _, prop := range ast.Properties {
		prop.Render(w)
	}
}

func (ast *DBEntry) ToHTML(w *html.Writer) {}
