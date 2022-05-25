package processors

import (
	"github.com/jomei/notionapi"
	"github.com/kapitanov/notion2html/internal/ast"
)

type headingProcessor struct {}

func (_ headingProcessor) Process(container ast.Container, provider Provider, rawBlock notionapi.Block) (ast.Node, error) {
	var level int
	var richText []notionapi.RichText

	switch b := rawBlock.(type) {
	case *notionapi.Heading1Block:
		level = 1
		richText = b.Heading1.RichText
	case *notionapi.Heading2Block:
		level = 2
		richText = b.Heading2.RichText
	case *notionapi.Heading3Block:
		level = 3
		richText = b.Heading3.RichText
	default:
		return nil, nil
	}

	ast := &ast.Heading{
		Level: level,
		Text:  buildText(richText),
	}

	container.AppendNode(ast)
	return ast, nil
}
