package processors

import (
	"github.com/jomei/notionapi"
	"github.com/kapitanov/notion2html/internal/ast"
)

type paragraphProcessor struct{}

func (_ paragraphProcessor) Process(container ast.Container, provider Provider, rawBlock notionapi.Block) (ast.Node, error) {
	block := rawBlock.(*notionapi.ParagraphBlock)
	ast := &ast.Paragraph{
		Text: buildText(block.Paragraph.RichText),
	}

	if len(ast.Text.Runs) > 0 {
		container.AppendNode(ast)
		return ast, nil
	}

	return nil, nil
}
