package ast

import (
	"github.com/kapitanov/notion2html/internal/html"
)

type Table struct {
	HasColumnHeader bool
	HasRowHeader    bool
	Rows            []*TableRow
	HeaderRow       *TableRow
	BodyRows        []*TableRow
}

func (ast *Table) NewRow() *TableRow {
	return ast.AddRow(&TableRow{})
}

func (ast *Table) AddRow(row *TableRow) *TableRow {
	ast.Rows = append(ast.Rows, row)
	return row
}

func (ast *Table) ToHTML(w *html.Writer) {
	if len(ast.Rows) == 0 {
		return
	}

	if ast.HasColumnHeader {
		ast.HeaderRow = ast.Rows[0]
		ast.BodyRows = ast.Rows[1:]
	} else {
		ast.HeaderRow = nil
		ast.BodyRows = ast.Rows
	}

	for _, row := range ast.Rows {
		if ast.HasRowHeader {
			row.HeaderCell = row.Cells[0]
			row.BodyCells = row.Cells[1:]
		} else {
			row.HeaderCell = nil
			row.BodyCells = row.Cells
		}

		for _, cell := range row.Cells {
			cell.Content = w.Render(func(wr *html.Writer) error {
				cell.Text.ToHTML(wr)
				return nil
			})
		}
	}

	w.Template("table.html", ast)
}

func (ast *Table) GetNodes() []Node {
	return nil
}

func (ast *Table) AppendNode(node Node) {}

func (ast *Table) ShouldProcessChildren() bool {
	return false
}

type TableRow struct {
	Cells []*TableCell

	HeaderCell *TableCell
	BodyCells  []*TableCell
}

func (ast *TableRow) NewCell() *TableCell {
	return ast.AddCell(&TableCell{})
}

func (ast *TableRow) AddCell(row *TableCell) *TableCell {
	ast.Cells = append(ast.Cells, row)
	return row
}

type TableCell struct {
	Text    *Text
	Content string
}

func (ast *TableCell) WithText(text *Text) *TableCell {
	ast.Text = text
	return ast
}
