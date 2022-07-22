package processors

import (
	"github.com/jomei/notionapi"

	"github.com/kapitanov/notion2html/internal/ast"
)

type videoProcessor struct{}

func (_ videoProcessor) Process(container ast.Container, provider Provider, rawBlock notionapi.Block) (ast.Node, error) {
	block := rawBlock.(*notionapi.VideoBlock)
	ast := &ast.Video{
		Caption: buildText(block.Video.Caption),
		URL:     block.Video.External.URL,
	}

	container.AppendNode(ast)
	return ast, nil
}
