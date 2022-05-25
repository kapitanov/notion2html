package tree

import (
	"context"

	"github.com/jomei/notionapi"
)

func Load(ctx context.Context, notion *notionapi.Client) (*PageSet, error) {
	rawPages, err := fetchPages(ctx, notion)
	if err != nil {
		return nil, err
	}

	rootPages, _ := buildTree(rawPages)
	setPagesDepth(rootPages, 0)

	set := &PageSet{
		Roots:       rootPages,
		LastUpdated: rawPages[0].LastEditedTime,
	}

	for _, page := range rawPages {
		if set.LastUpdated.Before(page.LastEditedTime) {
			set.LastUpdated = page.LastEditedTime
		}
	}

	return set, nil
}

func buildTree(rawPages []*notionapi.Page) (Pages, Pages) {
	idToRawPage := make(map[string]*notionapi.Page)
	idToPage := make(map[string]*Page)
	unreferencedPages := make(map[string]struct{})

	for _, rawPage := range rawPages {
		id := string(rawPage.ID)
		idToRawPage[id] = rawPage
		idToPage[id] = newPage(rawPage)
		unreferencedPages[id] = struct{}{}
	}

	for _, rawPage := range rawPages {
		pageID := string(rawPage.ID)

		page := idToPage[pageID]
		parentPage, ok := idToPage[string(rawPage.Parent.PageID)]
		if ok {
			parentPage.Children = append(parentPage.Children, page)
			page.Parent = parentPage
			delete(unreferencedPages, pageID)
		}
	}

	var rootPages []*Page
	for _, rawPage := range rawPages {
		if rawPage.Parent.Type == "workspace" {
			id := string(rawPage.ID)
			rootPages = append(rootPages, idToPage[id])
			delete(unreferencedPages, id)
		}
	}

	var freePages []*Page
	for _, rawPage := range rawPages {
		id := string(rawPage.ID)
		_, ok := unreferencedPages[id]
		if ok {
			freePages = append(freePages, idToPage[id])
		}
	}

	return rootPages, freePages
}

func setPagesDepth(pages Pages, depth int) {
	for _, page := range pages {
		page.Depth = depth
		setPagesDepth(page.Children, depth+1)
	}
}
