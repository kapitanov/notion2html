package tree

import (
	"context"

	"github.com/jomei/notionapi"
)

func fetchPages(ctx context.Context, notion *notionapi.Client) ([]*notionapi.Page, error) {
	req := &notionapi.SearchRequest{
		Filter: map[string]interface{}{
			"property": "object",
			"value":    "page",
		},
		PageSize: 100,
	}

	var pages []*notionapi.Page
	for {
		resp, err := notion.Search.Do(ctx, req)
		if err != nil {
			return nil, err
		}

		for _, result := range resp.Results {
			page, ok := result.(*notionapi.Page)
			if ok {
				//if page.Parent.Type == "page_id" || page.Parent.Type == "workspace" {
					pages = append(pages, page)
				//}
			}
		}

		if !resp.HasMore {
			break
		}

		req.StartCursor = resp.NextCursor
	}

	return pages, nil
}
