package processors

import (
	"github.com/jomei/notionapi"
	"github.com/kapitanov/notion2html/internal/ast"
)

type listProcessor struct{}

func (p listProcessor) Process(container ast.Container, provider Provider, rawBlock notionapi.Block) (ast.Node, error) {
	switch b := rawBlock.(type) {
	case *notionapi.BulletedListItemBlock:
		return p.processBulletList(container, provider, b)
	case *notionapi.NumberedListItemBlock:
		return p.processNumberedList(container, provider, b)
	case *notionapi.ToDoBlock:
		return p.processTodoList(container, provider, b)

	default:
		return nil, nil
	}
}

func (p listProcessor) processBulletList(
	container ast.Container,
	provider Provider,
	block *notionapi.BulletedListItemBlock,
) (ast.Node, error) {
	item := &ast.ListItem{
		Text: buildText(block.BulletedListItem.RichText),
	}

	for _, child := range block.BulletedListItem.Children {
		err := Process(item, provider, child)
		if err != nil {
			return nil, err
		}
	}

	var list *ast.BulletList
	nodes := container.GetNodes()

	if len(nodes) > 0 {
		var ok bool
		list, ok = nodes[len(nodes)-1].(*ast.BulletList)
		if !ok {
			list = nil
		}
	}

	if list == nil {
		list = &ast.BulletList{}
		container.AppendNode(list)
	}

	list.Items = append(list.Items, item)
	return item, nil
}

func (p listProcessor) processNumberedList(
	container ast.Container,
	provider Provider,
	block *notionapi.NumberedListItemBlock,
) (ast.Node, error) {
	item := &ast.ListItem{
		Text: buildText(block.NumberedListItem.RichText),
	}

	for _, child := range block.NumberedListItem.Children {
		err := Process(item, provider, child)
		if err != nil {
			return nil, err
		}
	}

	var list *ast.NumberedList
	nodes := container.GetNodes()

	if len(nodes) > 0 {
		var ok bool
		list, ok = nodes[len(nodes)-1].(*ast.NumberedList)
		if !ok {
			list = nil
		}
	}

	if list == nil {
		list = &ast.NumberedList{}
		container.AppendNode(list)
	}

	list.Items = append(list.Items, item)
	return item, nil
}

func (p listProcessor) processTodoList(
	container ast.Container,
	provider Provider,
	block *notionapi.ToDoBlock,
) (ast.Node, error) {
	item := &ast.ListItem{
		Text:    buildText(block.ToDo.RichText),
		Checked: block.ToDo.Checked,
	}

	for _, child := range block.ToDo.Children {
		err := Process(item, provider, child)
		if err != nil {
			return nil, err
		}
	}

	var list *ast.TodoList
	nodes := container.GetNodes()

	if len(nodes) > 0 {
		var ok bool
		list, ok = nodes[len(nodes)-1].(*ast.TodoList)
		if !ok {
			list = nil
		}
	}

	if list == nil {
		list = &ast.TodoList{}
		container.AppendNode(list)
	}

	list.Items = append(list.Items, item)
	return item, nil
}
