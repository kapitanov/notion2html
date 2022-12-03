package ast

import (
	"fmt"

	"github.com/kapitanov/notion2html/internal/html"
	"github.com/kapitanov/notion2html/internal/tree"
)

type Breadcrumb struct {
	Page  *tree.Page
	Items []*BreadcrumbItem
}

type BreadcrumbItem struct {
	Title string
	URL   string
}

func (ast *Breadcrumb) ToHTML(w *html.Writer) {
	ast.Items = ast.Items[0:0]

	page := ast.Page
	for page != nil {
		ast.Items = append(ast.Items, &BreadcrumbItem{
			Title: page.Title,
			URL:   fmt.Sprintf("%s.html", page.ID),
		})
		page = page.Parent
	}

	ast.Items[0].URL = ""

	for i := 0; i < len(ast.Items)/2; i++ {
		j := len(ast.Items) - i - 1
		ast.Items[i], ast.Items[j] = ast.Items[j], ast.Items[i]
	}

	w.Template("breadcrumb.html", ast)
}
