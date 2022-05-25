package ast

import (
	"github.com/kapitanov/notion2html/internal/html"
)

type Callout struct {
	Title   string
	Text    *Text
	Content string
	Nodes   []Node
}

func (ast *Callout) ToHTML(w *html.Writer) {
	ast.Content = w.Render(func(wr *html.Writer) error {
		ast.Text.ToHTML(wr)
		for _, node := range ast.Nodes {
			node.ToHTML(wr)
		}
		return nil
	})
	w.Template("callout.html", ast)
}

func (ast *Callout) GetNodes() []Node {
	return ast.Nodes
}

func (ast *Callout) AppendNode(node Node) {
	ast.Nodes = append(ast.Nodes, node)
}
