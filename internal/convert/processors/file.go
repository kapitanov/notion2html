package processors

import (
	"github.com/jomei/notionapi"
	"github.com/kapitanov/notion2html/internal/ast"
)

type fileProcessor struct{}

func (_ fileProcessor) Process(container ast.Container, provider Provider, rawBlock notionapi.Block) (ast.Node, error) {
	block := rawBlock.(*notionapi.FileBlock)
	ast := &ast.File{}

	if block.File.External != nil {
		ast.ExternalHref = block.File.External.URL
	}

	container.AppendNode(ast)
	return ast, nil
}
