package processors

import (
	"github.com/jomei/notionapi"
	"github.com/kapitanov/notion2html/internal/ast"
)

type toggleProcessor struct{}

func (_ toggleProcessor) Process(container ast.Container, provider Provider, rawBlock notionapi.Block) (ast.Node, error) {
	block := rawBlock.(*notionapi.ToggleBlock)
	toggle := &ast.Toggle{
		ID:   string(rawBlock.GetID()),
		Text: buildText(block.Toggle.RichText),
	}

	children, err := provider.ProvideChildren(block)
	if err != nil {
		return nil, err
	}

	for _, child := range children {
		err = Process(toggle, provider, child)
		if err != nil {
			return nil, err
		}
	}

	container.AppendNode(toggle)
	return toggle, nil
}
