package ast

import (
	"github.com/kapitanov/notion2html/internal/html"
)

type DBEntryContent struct {
	Nodes []Node
	HTML  string
}

func (ast *DBEntryContent) GetNodes() []Node {
	return ast.Nodes
}

func (ast *DBEntryContent) AppendNode(node Node) {
	ast.Nodes = append(ast.Nodes, node)
}

func (ast *DBEntryContent) ToHTML(w *html.Writer) {}

func (ast *DBEntryContent) Render(w *html.Writer) {
	ast.HTML = w.Render(func(wr *html.Writer) error {
		for _, node := range ast.Nodes {
			node.ToHTML(wr)
		}
		return nil
	})
}
