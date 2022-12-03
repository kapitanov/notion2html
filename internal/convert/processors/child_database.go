package processors

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/jomei/notionapi"
	"github.com/kapitanov/notion2html/internal/ast"
)

type childDatabaseProcessor struct{}

func (c childDatabaseProcessor) Process(container ast.Container, provider Provider, rawBlock notionapi.Block) (ast.Node, error) {
	db, err := provider.ProvideDatabase(rawBlock)
	if err != nil {
		return nil, err
	}

	data, err := provider.ProvideDatabaseData(db)
	if err != nil {
		return nil, err
	}

	node := &ast.DB{
		ID: string(db.ID),
	}
	err = c.processData(provider, node, db, data)
	if err != nil {
		return nil, err
	}

	container.AppendNode(node)
	return node, nil
}

func (c childDatabaseProcessor) processData(provider Provider, db *ast.DB, rawDB *notionapi.Database, rawItems []notionapi.Page) error {
	var properties []struct {
		Name string
		Type notionapi.PropertyConfigType
	}

	for name, prop := range rawDB.Properties {
		properties = append(properties, struct {
			Name string
			Type notionapi.PropertyConfigType
		}{
			Name: name,
			Type: prop.GetType(),
		})
	}

	var less = func(i, j int) bool {
		if properties[i].Type == notionapi.PropertyConfigTypeTitle {
			if properties[j].Type != notionapi.PropertyConfigTypeTitle {
				return true
			}
		}

		if properties[j].Type == notionapi.PropertyConfigTypeTitle {
			if properties[i].Type != notionapi.PropertyConfigTypeTitle {
				return false
			}
		}

		return strings.Compare(properties[i].Name, properties[j].Name) < 0
	}

	sort.Slice(properties, less)

	for _, prop := range properties {
		db.PropertyNames = append(db.PropertyNames, prop.Name)
	}

	for _, rawItem := range rawItems {
		err := c.processDataItem(provider, db.NewEntry(), rawItem, db.PropertyNames)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c childDatabaseProcessor) processDataItem(provider Provider, entry *ast.DBEntry, rawItem notionapi.Page, propertyNames []string) error {
	entry.ID = string(rawItem.ID)

	for _, propName := range propertyNames {
		propValue, ok := rawItem.Properties[propName]
		if !ok {
			continue
		}

		entry.NewProperty(propName, c.processDataItemProperty(propName, propValue))
	}

	children, err := provider.ProvidePageChildren(rawItem)
	if err != nil {
		return err
	}

	content := &ast.DBEntryContent{}
	err = c.processDataItemContent(content, provider, children)
	if err != nil {
		return err
	}

	if len(content.Nodes) > 0 {
		entry.Content = content
	}

	return nil
}

func (c childDatabaseProcessor) processDataItemContent(container ast.Container, provider Provider, blocks []notionapi.Block) error {
	for _, block := range blocks {
		err := Process(container, provider, block)
		if err != nil {
			return err
		}
	}

	return nil
}

func (_ childDatabaseProcessor) processDataItemProperty(name string, value notionapi.Property) ast.DBEntryPropertyValue {
	switch p := value.(type) {

	case *notionapi.TitleProperty:
		return ast.NewDBEntryPropertyTitleValue(buildRawText(p.Title))

	case *notionapi.RichTextProperty:
		return ast.NewDBEntryPropertyRichTextValue(buildText(p.RichText))

	case *notionapi.TextProperty:
		return ast.NewDBEntryPropertyTextValue(buildRawText(p.Text))

	case *notionapi.NumberProperty:
		return ast.NewDBEntryPropertyTextValue(fmt.Sprintf("%v", p.Number))

	case *notionapi.SelectProperty:
		return ast.NewDBEntryPropertyTagsValue(
			[]*ast.DBEntryTag{
				{
					Name:  p.Select.Name,
					Color: string(p.Select.Color),
				},
			})

	case *notionapi.MultiSelectProperty:
		array := make([]*ast.DBEntryTag, len(p.MultiSelect))
		for i := range p.MultiSelect {
			array[i] = &ast.DBEntryTag{
				Name:  p.MultiSelect[i].Name,
				Color: string(p.MultiSelect[i].Color),
			}
		}
		return ast.NewDBEntryPropertyTagsValue(array)

	case *notionapi.DateProperty:
		var dateStr string
		if p.Date.Start != nil && p.Date.End != nil {
			dateStr = fmt.Sprintf("%s - %s", p.Date.Start, p.Date.End)
		} else if p.Date.Start != nil {
			dateStr = fmt.Sprintf("%s", p.Date.Start)
		} else if p.Date.End != nil {
			dateStr = fmt.Sprintf("%s", p.Date.End)
		} else {
			dateStr = ""
		}
		return ast.NewDBEntryPropertyTextValue(dateStr)

	case *notionapi.CheckboxProperty:
		return ast.NewDBEntryPropertyBoolValue(p.Checkbox)

	case *notionapi.URLProperty:
		return ast.NewDBEntryPropertyRichTextValue(ast.NewHrefText(p.URL, p.URL))

	case *notionapi.EmailProperty:
		return ast.NewDBEntryPropertyRichTextValue(ast.NewHrefText(p.Email, "mailto:"+p.Email))

	case *notionapi.PhoneNumberProperty:
		return ast.NewDBEntryPropertyRichTextValue(ast.NewHrefText(p.PhoneNumber, "tel:"+p.PhoneNumber))

	default:
		log.Printf("unknown property '%s'", name)
		return ast.NewDBEntryPropertyRichTextValue(ast.NewPlainText(fmt.Sprintf("unknown property '%s'", name)))
	}
}
