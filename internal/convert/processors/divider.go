package processors

import (
	"github.com/jomei/notionapi"
	"github.com/kapitanov/notion2html/internal/ast"
)

type dividerProcessor struct{}

func (_ dividerProcessor) Process(container ast.Container, provider Provider, rawBlock notionapi.Block) (ast.Node, error) {
	ast := &ast.Divider{}

	container.AppendNode(ast)
	return ast, nil
}
