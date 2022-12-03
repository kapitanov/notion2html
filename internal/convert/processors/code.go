package processors

import (
	"github.com/jomei/notionapi"
	"github.com/kapitanov/notion2html/internal/ast"
)

type codeProcessor struct{}

func (_ codeProcessor) Process(container ast.Container, provider Provider, rawBlock notionapi.Block) (ast.Node, error) {
	block := rawBlock.(*notionapi.CodeBlock)
	ast := &ast.Code{
		Language: block.Code.Language,
	}

	for _, rt := range block.Code.RichText {
		ast.Lines = append(ast.Lines, rt.PlainText)
	}

	container.AppendNode(ast)
	return ast, nil
}
