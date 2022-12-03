package ast

import (
	"strings"

	"github.com/kapitanov/notion2html/internal/html"
)

type Code struct {
	Language string
	Lines    []string
	Text     string
}

func (ast *Code) ToHTML(w *html.Writer) {
	ast.Text = strings.Join(ast.Lines, "\n")
	w.Template("code.html", ast)
}
