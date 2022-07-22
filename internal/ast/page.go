package ast

import (
	"github.com/kapitanov/notion2html/internal/html"
	"github.com/kapitanov/notion2html/internal/tree"
)

type Page struct {
	Page    *tree.Page
	Title   string
	Content string
	Breadcrumb string
	Nodes   []Node
}

func (ast *Page) ToHTML(w *html.Writer) {
	ast.Breadcrumb = w.Render(func(wr *html.Writer) error {
		breadcrumb := &Breadcrumb{
			Page: ast.Page,
		}
		breadcrumb.ToHTML(wr)
		return nil
	})

	ast.Content = w.Render(func(wr *html.Writer) error {
		for _, node := range ast.Nodes {
			node.ToHTML(wr)
		}
		return nil
	})

	w.Template("page.html", ast)
}

func (ast *Page) GetNodes() []Node {
	return ast.Nodes
}

func (ast *Page) AppendNode(node Node) {
	ast.Nodes = append(ast.Nodes, node)
}
