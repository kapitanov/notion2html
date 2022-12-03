package processors

import (
	"github.com/jomei/notionapi"
	"github.com/kapitanov/notion2html/internal/ast"
)

type calloutProcessor struct{}

func (_ calloutProcessor) Process(container ast.Container, provider Provider, rawBlock notionapi.Block) (ast.Node, error) {
	block := rawBlock.(*notionapi.CalloutBlock)
	callout := &ast.Callout{
		Text: buildText(block.Callout.RichText),
	}

	if block.Callout.Icon != nil {
		if block.Callout.Icon.Emoji != nil {
			callout.Title = string(*block.Callout.Icon.Emoji)
		}
	}

	children, err := provider.ProvideChildren(block)
	if err != nil {
		return nil, err
	}

	for _, child := range children {
		err = Process(callout, provider, child)
		if err != nil {
			return nil, err
		}
	}

	container.AppendNode(callout)
	return callout, nil
}
