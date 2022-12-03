package processors

import (
	"github.com/jomei/notionapi"
	"github.com/kapitanov/notion2html/internal/ast"
)

type quoteProcessor struct{}

func (_ quoteProcessor) Process(container ast.Container, provider Provider, rawBlock notionapi.Block) (ast.Node, error) {
	block := rawBlock.(*notionapi.QuoteBlock)
	ast := &ast.Quote{
		Text: buildText(block.Quote.RichText),
	}

	container.AppendNode(ast)
	return ast, nil
}
