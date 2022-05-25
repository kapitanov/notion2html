package processors

import (
	"log"

	"github.com/jomei/notionapi"
	"github.com/kapitanov/notion2html/internal/ast"
)

func wrapChildrenProcessor(processor Processor) Processor {
	return &childrenProcessorWrapper{processor}
}

type childrenProcessorWrapper struct {
	processor Processor
}

func (p *childrenProcessorWrapper) Process(
	container ast.Container,
	provider Provider,
	rawBlock notionapi.Block,
) (ast.Node, error) {
	node, err := p.processor.Process(container, provider, rawBlock)
	if err != nil {
		return nil, err
	}

	if rawBlock.GetHasChildren() {
		if node != nil {
			nestedContainer, ok := node.(ast.Container)
			if !ok {
				log.Printf("block \"%s\" has children but its parent is not a container", rawBlock.GetType())
			} else {
				nestedContainerEx, ok := node.(ast.ContainerEx)
				if !ok || !nestedContainerEx.ShouldProcessChildren() {
					children, err := provider.ProvideChildren(rawBlock)
					if err != nil {
						return nil, err
					}

					for _, child := range children {
						err = Process(nestedContainer, provider, child)
					}
				}
			}
		}
	}

	return node, nil
}
