package convert

import (
	"context"
	"fmt"

	"github.com/jomei/notionapi"
	"github.com/kapitanov/notion2html/internal/ast"
	"github.com/kapitanov/notion2html/internal/convert/processors"
	"github.com/kapitanov/notion2html/internal/tree"
)

type ASTBuilder struct {
	notion *notionapi.Client
}

func NewASTBuilder(notion *notionapi.Client) *ASTBuilder {
	return &ASTBuilder{
		notion: notion,
	}
}

func (b *ASTBuilder) Build(ctx context.Context, page *tree.Page) (*ast.Page, error) {
	blocks, err := b.getBlocks(ctx, notionapi.BlockID(page.ID))
	if err != nil {
		return nil, err
	}

	ast, err := b.buildPage(page, blocks)
	if err != nil {
		return nil, err
	}

	return ast, nil
}

func (b *ASTBuilder) getTitle(page *notionapi.Page) string {
	prop, ok := page.Properties["title"]
	if ok {
		titleProp, ok := prop.(*notionapi.TitleProperty)
		if ok {
			str := ""
			for _, t := range titleProp.Title {
				str += t.PlainText
			}
			return str
		}
	}

	return string(page.ID)
}

func (b *ASTBuilder) getBlocks(ctx context.Context, blockID notionapi.BlockID) ([]notionapi.Block, error) {
	pagination := &notionapi.Pagination{
		PageSize: 100,
	}

	var blocks []notionapi.Block
	for {
		resp, err := b.notion.Block.GetChildren(ctx, blockID, pagination)
		if err != nil {
			return nil, err
		}

		blocks = append(blocks, resp.Results...)

		if !resp.HasMore {
			break
		}

		pagination.StartCursor = notionapi.Cursor(resp.NextCursor)
	}

	return blocks, nil
}

func (b *ASTBuilder) buildPage(page *tree.Page, blocks []notionapi.Block) (*ast.Page, error) {
	pageAST := &ast.Page{
		Page:  page,
		Title: page.Title,
	}

	if len(page.Children) > 0 {
		childList := &ast.ChildList{}
		pageAST.AppendNode(childList)

		for _, child := range page.Children {
			item := childList.NewItem()
			item.Title = child.Title
			item.URL = fmt.Sprintf("%s.html", child.ID)
		}
	}

	for _, block := range blocks {
		err := processors.Process(pageAST, b, block)
		if err != nil {
			return nil, err
		}
	}

	return pageAST, nil
}

func (b *ASTBuilder) ProvideChildren(rawBlock notionapi.Block) ([]notionapi.Block, error) {
	return b.getBlocks(context.Background(), rawBlock.GetID())
}

func (b *ASTBuilder) ProvidePageChildren(rawBlock notionapi.Page) ([]notionapi.Block, error) {
	return b.getBlocks(context.Background(), notionapi.BlockID(rawBlock.ID.String()))
}

func (b *ASTBuilder) ProvideDatabase(rawBlock notionapi.Block) (*notionapi.Database, error) {
	resp, err := b.notion.Database.Get(context.Background(), notionapi.DatabaseID(rawBlock.GetID()))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *ASTBuilder) ProvideDatabaseData(db *notionapi.Database) ([]notionapi.Page, error) {
	ctx := context.Background()
	req := &notionapi.DatabaseQueryRequest{
		PageSize: 100,
	}

	var pages []notionapi.Page
	for {
		resp, err := b.notion.Database.Query(ctx, notionapi.DatabaseID(db.ID), req)
		if err != nil {
			return nil, err
		}

		pages = append(pages, resp.Results...)

		if !resp.HasMore {
			break
		}

		req.StartCursor = notionapi.Cursor(resp.NextCursor)
	}

	return pages, nil
}
