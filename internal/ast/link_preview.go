package ast

import (
	"github.com/kapitanov/notion2html/internal/html"
	"net/url"
)

type LinkPreview struct {
	URL   string
	Title string
}

func (ast *LinkPreview) ToHTML(w *html.Writer) {
	u, err := url.Parse(ast.URL)
	if err != nil {
		ast.Title = ast.URL
	} else {
		ast.Title = u.Host
	}

	w.Template("link_preview.html", ast)
}
