package ast

import (
	"github.com/kapitanov/notion2html/internal/html"
)

type DB struct {
	ID            string
	PropertyNames []string
	Entries       []*DBEntry
}

func (ast *DB) NewEntry() *DBEntry {
	entry := &DBEntry{}
	ast.Entries = append(ast.Entries, entry)
	return entry
}

func (ast *DB) ToHTML(w *html.Writer) {
	if len(ast.Entries) == 0 {
		return
	}

	for _, entry := range ast.Entries {
		entry.Render(w)
	}

	w.Template("db.html", ast)
}
