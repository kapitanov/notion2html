package ast

import (
	"path/filepath"

	"github.com/kapitanov/notion2html/internal/html"
)

type File struct {
	ExternalHref string
	Name         string
}

func (ast *File) ToHTML(w *html.Writer) {
	_, name := filepath.Split(ast.ExternalHref)
	ast.Name = name
}
