package ast

import (
	"github.com/kapitanov/notion2html/internal/html"
)

type Image struct {
	Href       string
	CachedHref string
}

func (ast *Image) ToHTML(w *html.Writer) {
	ast.CachedHref = w.CacheURL(ast.Href)
	w.Template("image.html", ast)
}
