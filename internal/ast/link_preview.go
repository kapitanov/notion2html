package ast

import (
	"github.com/kapitanov/notion2html/internal/html"
)

type LinkPreview struct {
	URL         string
	Title       string
	Icon        string
	Name        string
	Description string
	Images      []string
	Link        string
}

func (ast *LinkPreview) ToHTML(w *html.Writer) {
	w.Template("link_preview.html", ast)
}
