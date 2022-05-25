package processors

import (
	"github.com/jomei/notionapi"
	"github.com/kapitanov/notion2html/internal/ast"
)

type linkPreviewProcessor struct{}

func (_ linkPreviewProcessor) Process(container ast.Container, provider Provider, rawBlock notionapi.Block) (ast.Node, error) {
	block := rawBlock.(*notionapi.LinkPreviewBlock)
	linkPreview := &ast.LinkPreview{
		URL: block.LinkPreview.URL,
	}

	container.AppendNode(linkPreview)
	return linkPreview, nil
}
