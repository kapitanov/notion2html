package ast

import (
	"github.com/kapitanov/notion2html/internal/html"
)

type Quote struct {
	Text    *Text
	Nodes   []Node
	Content string
}

func (ast *Quote) ToHTML(w *html.Writer) {
	ast.Content = w.Render(func(wr *html.Writer) error {
		ast.Text.ToHTML(wr)
		for _, node := range ast.Nodes {
			node.ToHTML(wr)
		}
		return nil
	})

	w.Template("quote.html", ast)
}

func (ast *Quote) GetNodes() []Node {
	return ast.Nodes
}

func (ast *Quote) AppendNode(node Node) {
	ast.Nodes = append(ast.Nodes, node)
}
