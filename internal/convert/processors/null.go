package processors

import (
	"log"

	"github.com/jomei/notionapi"
	"github.com/kapitanov/notion2html/internal/ast"
)

type nullBlockProcessor struct{}

func (_ nullBlockProcessor) Process(
	container ast.Container,
	provider Provider,
	rawBlock notionapi.Block,
) (ast.Node, error) {
	return nil, nil
}

var unsupportedBlockProcessor Processor = unsupportedBlockProcessorStub{}

type unsupportedBlockProcessorStub struct{}

func (_ unsupportedBlockProcessorStub) Process(
	container ast.Container,
	provider Provider,
	rawBlock notionapi.Block,
) (ast.Node, error) {
	log.Printf("block \"%s\" is not supported", rawBlock.GetType())
	return nil, nil
}
