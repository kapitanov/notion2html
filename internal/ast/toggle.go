package ast

import (
	"github.com/kapitanov/notion2html/internal/html"
)

type Toggle struct {
	ID      string
	Text    *Text
	Nodes   []Node
	Header  string
	Content string
}

func (ast *Toggle) ToHTML(w *html.Writer) {
	ast.Header = w.Render(func(wr *html.Writer) error {
		ast.Text.ToHTML(wr)
		return nil
	})

	ast.Content = w.Render(func(wr *html.Writer) error {
		for _, node := range ast.Nodes {
			node.ToHTML(wr)
		}
		return nil
	})

	w.Template("toggle.html", ast)
	return
}

func (ast *Toggle) GetNodes() []Node {
	return ast.Nodes
}

func (ast *Toggle) AppendNode(node Node) {
	ast.Nodes = append(ast.Nodes, node)
}
