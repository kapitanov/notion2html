package ast

import (
	"github.com/kapitanov/notion2html/internal/html"
)

type Heading struct {
	Level   int
	Text    *Text
	Content string
}

func (ast *Heading) ToHTML(w *html.Writer) {
	ast.Content = w.Render(func(wr *html.Writer) error {
		ast.Text.ToHTML(wr)
		return nil
	})

	switch ast.Level {
	case 1:
		w.Template("h1.html", ast)
	case 2:
		w.Template("h2.html", ast)
	case 3:
		w.Template("h3.html", ast)
	default:
		w.WithTag(html.Tagf("h%d", ast.Level+1), func() error {
			w.WriteLine(ast.Content)
			return nil
		})
	}
}
