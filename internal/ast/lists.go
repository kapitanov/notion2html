package ast

import (
	"github.com/kapitanov/notion2html/internal/html"
)

type NumberedList struct {
	Items []*ListItem
}

func (ast *NumberedList) ToHTML(w *html.Writer) {
	for _, item := range ast.Items {
		item.Render(w)
	}

	w.Template("ol.html", ast)
}

type BulletList struct {
	Items []*ListItem
}

func (ast *BulletList) ToHTML(w *html.Writer) {
	for _, item := range ast.Items {
		item.Render(w)
	}

	w.Template("ul.html", ast)
}

type TodoList struct {
	Items []*ListItem
}

func (ast *TodoList) ToHTML(w *html.Writer) {
	for _, item := range ast.Items {
		item.Render(w)
	}

	w.Template("todo.html", ast)
}

type ListItem struct {
	Text    *Text
	Nodes   []Node
	Checked bool
	Content string
}

func (ast *ListItem) Render(w *html.Writer) {
	ast.Content = w.Render(func(wr *html.Writer) error {
		ast.Text.ToHTML(wr)
		for _, node := range ast.Nodes {
			node.ToHTML(wr)
		}
		return nil
	})
}

func (ast *ListItem) ToHTML(w *html.Writer) {
}

func (ast *ListItem) GetNodes() []Node {
	return ast.Nodes
}

func (ast *ListItem) AppendNode(node Node) {
	ast.Nodes = append(ast.Nodes, node)
}
