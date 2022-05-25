package ast

import (
	"github.com/kapitanov/notion2html/internal/html"
)

type ChildList struct {
	Items []*ChildListItem
}

func (ast *ChildList) NewItem() *ChildListItem {
	return ast.AddItem(&ChildListItem{})
}

func (ast *ChildList) AddItem(row *ChildListItem) *ChildListItem {
	ast.Items = append(ast.Items, row)
	return row
}

func (ast *ChildList) ToHTML(w *html.Writer) {
	if len(ast.Items) == 0 {
		return
	}

	w.Template("child_list.html", ast)
}

type ChildListItem struct {
	Title string
	URL   string
}
