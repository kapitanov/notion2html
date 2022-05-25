package ast

import (
	"github.com/kapitanov/notion2html/internal/html"
)

type Divider struct {
}

func (ast *Divider) ToHTML(w *html.Writer) {
	w.Template("divider.html", nil)
}
