package ast

import (
	"github.com/kapitanov/notion2html/internal/html"
)

type Paragraph struct {
	Text    *Text
	Nodes   []Node
	Content string
}

func (ast *Paragraph) ToHTML(w *html.Writer) {
	ast.Content = w.Render(func(wr *html.Writer) error {
		ast.Text.ToHTML(wr)
		for _, node := range ast.Nodes {
			node.ToHTML(wr)
		}
		return nil
	})

	w.Template("paragraph.html", ast)
}

func (ast *Paragraph) GetNodes() []Node {
	return ast.Nodes
}

func (ast *Paragraph) AppendNode(node Node) {
	ast.Nodes = append(ast.Nodes, node)
}
