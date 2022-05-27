package tree

import (
	"time"

	"github.com/jomei/notionapi"
)

type PageSet struct {
	Roots       Pages
	ByID        map[string]*Page
	LastUpdated time.Time
}

type Pages []*Page

func (pages Pages) Traverse(fn func(page *Page) error) error {
	for _, page := range pages {
		err := page.Traverse(fn)
		if err != nil {
			return err
		}
	}
	return nil
}

type Page struct {
	ID         string
	Title      string
	LastEdited time.Time
	Children   Pages
	URL        string
	Parent     *Page
	Depth      int
}

func (p *Page) Traverse(fn func(page *Page) error) error {
	if p == nil {
		return nil
	}

	err := fn(p)
	if err != nil {
		return err
	}

	for _, child := range p.Children {
		err = child.Traverse(fn)
		if err != nil {
			return err
		}
	}

	return nil
}

func newPage(page *notionapi.Page) *Page {
	p := &Page{
		ID:         string(page.ID),
		Title:      getTitle(page),
		LastEdited: page.LastEditedTime,
		URL:        page.URL,
	}
	return p
}

func getTitle(page *notionapi.Page) string {
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
