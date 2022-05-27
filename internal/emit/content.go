package emit

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/kapitanov/notion2html/internal/ast"
	"github.com/kapitanov/notion2html/internal/html"
	"github.com/kapitanov/notion2html/internal/tree"
)

func (e *Emitter) Generate(ctx context.Context, pageSet *tree.PageSet) (int, error) {
	e.pageCount = 0
	err := pageSet.Roots.Traverse(func(page *tree.Page) error {
		return e.contentPage(ctx, page)
	})
	if err != nil {
		return 0, err
	}

	err = e.treeJSON(pageSet)
	if err != nil {
		return 0, err
	}

	err = e.indexPage(pageSet)
	if err != nil {
		return 0, err
	}

	err = e.dropRemovedPages(pageSet)
	if err != nil {
		return 0, err
	}

	err = e.metadata.Save()
	if err != nil {
		return 0, err
	}

	return 0, nil
}

func (e *Emitter) indexPage(pageSet *tree.PageSet) error {
	return e.emitHTML("index.html", func(w *html.Writer) error {
		pageAST := &ast.IndexPage{
			Pages:       pageSet.Roots,
			LastUpdated: pageSet.LastUpdated,
		}
		pageAST.ToHTML(w)
		return nil
	})
}

func (e *Emitter) contentPage(ctx context.Context, page *tree.Page) error {
	if e.metadata.IsUpToDate(page.ID, page.LastEdited) {
		return nil
	}

	log.Printf("generating page %v", page.ID)
	pageAST, err := e.builder.Build(ctx, page)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%s.html", page.ID)
	err = e.emitHTML(filename, func(w *html.Writer) error {
		pageAST.ToHTML(w)
		return nil
	})
	if err != nil {
		return err
	}

	e.metadata.UpdateLastEdited(page.ID, page.LastEdited)

	e.pageCount++
	return nil
}

func (e *Emitter) treeJSON(pageSet *tree.PageSet) error {
	json := NewPageTreeJSON(pageSet)

	err := e.emitJSON("index.json", json)
	if err != nil {
		return err
	}

	return nil
}

func (e *Emitter) dropRemovedPages(pageSet *tree.PageSet) error {
	entries, err := os.ReadDir(e.outputDirectory)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		id := filepath.Base(entry.Name())
		ext := filepath.Ext(id)
		if ext != ".html" {
			continue
		}

		id = id[0 : len(id)-len(ext)]
		if id == "index" {
			continue
		}

		if _, exists := pageSet.ByID[id]; exists {
			continue
		}

		log.Printf("page %s doesn't exists anymore", id)
		err = os.Remove(filepath.Join(e.outputDirectory, entry.Name()))
		if err != nil {
			return err
		}
	}

	return nil
}

type PageTreeJSON struct {
	Pages []*PageTreeItemJSON `json:"pages,omitempty"`
}

func NewPageTreeJSON(pageSet *tree.PageSet) *PageTreeJSON {
	json := &PageTreeJSON{}
	for _, page := range pageSet.Roots {
		json.Pages = append(json.Pages, NewPageTreeItemJSON(page))
	}
	return json
}

type PageTreeItemJSON struct {
	Title string              `json:"title"`
	URL   string              `json:"url"`
	Pages []*PageTreeItemJSON `json:"pages,omitempty"`
}

func NewPageTreeItemJSON(page *tree.Page) *PageTreeItemJSON {
	json := &PageTreeItemJSON{
		Title: page.Title,
		URL:   fmt.Sprintf("%s.html", page.ID),
	}

	if len(page.Children) > 0 {
		for _, page := range page.Children {
			json.Pages = append(json.Pages, NewPageTreeItemJSON(page))
		}
	}

	return json
}
