package processors

import (
	"github.com/jomei/notionapi"
	"github.com/kapitanov/notion2html/internal/ast"
)

type Provider interface {
	ProvideChildren(rawBlock notionapi.Block) ([]notionapi.Block, error)
	ProvidePageChildren(rawBlock notionapi.Page) ([]notionapi.Block, error)
	ProvideDatabase(rawBlock notionapi.Block) (*notionapi.Database, error)
	ProvideDatabaseData(db *notionapi.Database) ([]notionapi.Page, error)
}

type Processor interface {
	Process(container ast.Container, provider Provider, rawBlock notionapi.Block) (ast.Node, error)
}

type Func func(container ast.Container, provider Provider, rawBlock notionapi.Block) (ast.Node, error)

var blockProcessors = map[notionapi.BlockType]Processor{
	notionapi.BlockTypeHeading1:         headingProcessor{},
	notionapi.BlockTypeHeading2:         headingProcessor{},
	notionapi.BlockTypeHeading3:         headingProcessor{},
	notionapi.BlockTypeParagraph:        paragraphProcessor{},
	notionapi.BlockQuote:                quoteProcessor{},
	notionapi.BlockTypeImage:            imageProcessor{},
	notionapi.BlockTypeNumberedListItem: listProcessor{},
	notionapi.BlockTypeBulletedListItem: listProcessor{},
	notionapi.BlockTypeCode:             codeProcessor{},
	notionapi.BlockTypeDivider:          dividerProcessor{},
	notionapi.BlockTypeFile:             fileProcessor{},
	notionapi.BlockTypeChildPage:        nullBlockProcessor{},
	notionapi.BlockTypeChildDatabase:    childDatabaseProcessor{},
	notionapi.BlockTypeBreadcrumb:       nullBlockProcessor{},
	notionapi.BlockCallout:              calloutProcessor{},
	notionapi.BlockTypeTableBlock:       tableProcessor{},
	notionapi.BlockTypeTableOfContents:  nullBlockProcessor{},
	notionapi.BlockTypeTableRowBlock:    nullBlockProcessor{},
	notionapi.BlockTypeToggle:           toggleProcessor{},
	notionapi.BlockTypeToDo:             listProcessor{},
	notionapi.BlockTypeLinkPreview:      linkPreviewProcessor{},
}

func init() {
	for t := range blockProcessors {
		blockProcessors[t] = wrapChildrenProcessor(blockProcessors[t])
	}
}

func Process(container ast.Container, provider Provider, rawBlock notionapi.Block) error {
	_, err := For(rawBlock).Process(container, provider, rawBlock)
	return err
}

func For(block notionapi.Block) Processor {
	blockType := block.GetType()
	processor, ok := blockProcessors[blockType]
	if !ok {
		return unsupportedBlockProcessor
	}

	return processor
}

func buildText(richText []notionapi.RichText) *ast.Text {
	text := &ast.Text{}

	for _, textRun := range richText {
		element := &ast.TextRun{
			Bold:          textRun.Annotations.Bold,
			Italic:        textRun.Annotations.Italic,
			Strikethrough: textRun.Annotations.Strikethrough,
			Underline:     textRun.Annotations.Underline,
			Code:          textRun.Annotations.Code,
			Color:         string(textRun.Annotations.Color),
			Href:          textRun.Href,
			Text:          textRun.PlainText,
		}

		text.Runs = append(text.Runs, element)
	}

	return text
}

func buildRawText(richText []notionapi.RichText) string {
	text := ""

	for _, textRun := range richText {
		text += textRun.PlainText
	}

	return text
}
