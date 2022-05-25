package processors

import (
	"github.com/jomei/notionapi"
	"github.com/kapitanov/notion2html/internal/ast"
)

type imageProcessor struct{}

func (_ imageProcessor) Process(container ast.Container, provider Provider, rawBlock notionapi.Block) (ast.Node, error) {
	block := rawBlock.(*notionapi.ImageBlock)
	ast := &ast.Image{}

	if block.Image.External != nil {
		ast.Href = block.Image.External.URL
	}

	if block.Image.File != nil {
		ast.Href = block.Image.File.URL
	}

	container.AppendNode(ast)
	return ast, nil
}
