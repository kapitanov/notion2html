package ast

import (
	"time"

	"github.com/kapitanov/notion2html/internal/html"
	"github.com/kapitanov/notion2html/internal/tree"
)

type IndexPage struct {
	Pages          tree.Pages
	LastUpdated    time.Time
	Items          []*IndexPageItem
	LastUpdatedStr string
}

type IndexPageItem struct {
	Page  *tree.Page
	Depth int
}

func (ast *IndexPage) ToHTML(w *html.Writer) {
	ast.Items = ast.Items[0:0]
	_ = ast.Pages.Traverse(func(page *tree.Page) error {
		ast.Items = append(ast.Items, &IndexPageItem{
			Page:  page,
			Depth: page.Depth,
		})

		return nil
	})

	w.Template("index.html", ast)
}
