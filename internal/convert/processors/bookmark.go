package processors

import (
	"log"
	"net/url"

	"github.com/badoux/goscraper"
	"github.com/jomei/notionapi"

	"github.com/kapitanov/notion2html/internal/ast"
)

type bookmarkProcessor struct{}

func (_ bookmarkProcessor) Process(container ast.Container, provider Provider, rawBlock notionapi.Block) (ast.Node, error) {
	block := rawBlock.(*notionapi.BookmarkBlock)
	linkPreview := &ast.LinkPreview{
		URL: block.Bookmark.URL,
	}

	u, err := url.Parse(linkPreview.URL)
	if err != nil {
		linkPreview.Title = linkPreview.URL
	} else {
		linkPreview.Title = u.Host
	}

	document, err := goscraper.Scrape(linkPreview.URL, 5)
	if err != nil {
		log.Printf("unable to scrape %s: %v", linkPreview.URL, err)
	} else {
		linkPreview.Icon = document.Preview.Icon
		linkPreview.Name = document.Preview.Name
		linkPreview.Title = document.Preview.Title
		linkPreview.Description = document.Preview.Description
		linkPreview.Images = document.Preview.Images
		linkPreview.Link = document.Preview.Link
	}

	container.AppendNode(linkPreview)
	return linkPreview, nil
}
