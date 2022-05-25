package processors

import (
	"github.com/jomei/notionapi"
	"github.com/kapitanov/notion2html/internal/ast"
)

type tableProcessor struct{}

func (_ tableProcessor) Process(container ast.Container, provider Provider, rawBlock notionapi.Block) (ast.Node, error) {
	block := rawBlock.(*notionapi.TableBlock)
	table := &ast.Table{
		HasColumnHeader: block.Table.HasColumnHeader,
		HasRowHeader:    block.Table.HasRowHeader,
	}

	children, err := provider.ProvideChildren(block)
	if err != nil {
		return nil, err
	}

	for _, childRaw := range children {
		child, ok := childRaw.(*notionapi.TableRowBlock)
		if !ok {
			continue
		}

		row := table.NewRow()
		for _, c := range child.TableRow.Cells {
			cell := &ast.TableCell{
				Text: buildText(c),
			}
			row.AddCell(cell)
		}
	}

	container.AppendNode(table)
	return table, nil
}
