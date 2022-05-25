package ast

import (
	"fmt"

	"github.com/kapitanov/notion2html/internal/html"
)

type DBEntryProperty struct {
	Name    string
	Value   DBEntryPropertyValue
	Content string
}

func (v *DBEntryProperty) Render(w *html.Writer) {
	v.Content = ""

	if v.Value != nil {
		v.Content = w.Render(func(wr *html.Writer) error {
			v.Value.ToHTML(wr)
			return nil
		})
	}
}

func (v *DBEntryProperty) ToHTML(w *html.Writer) {}

type DBEntryPropertyValue interface {
	ToHTML(w *html.Writer)
}

type DBEntryPropertyTitleValue struct {
	Value string
}

func NewDBEntryPropertyTitleValue(value string) *DBEntryPropertyTitleValue {
	return &DBEntryPropertyTitleValue{
		Value: value,
	}
}

func (v *DBEntryPropertyTitleValue) ToHTML(w *html.Writer) {
	w.PushTag(html.Tag("strong"))
	w.WriteLine(html.Escape(v.Value))
	w.PopTag()
}

type DBEntryPropertyTextValue struct {
	Value string
}

func NewDBEntryPropertyTextValue(value string) *DBEntryPropertyTextValue {
	return &DBEntryPropertyTextValue{
		Value: value,
	}
}

func (v *DBEntryPropertyTextValue) ToHTML(w *html.Writer) {
	w.WriteLine(html.Escape(v.Value))
}

type DBEntryPropertyRichTextValue struct {
	Value *Text
}

func NewDBEntryPropertyRichTextValue(value *Text) *DBEntryPropertyRichTextValue {
	return &DBEntryPropertyRichTextValue{
		Value: value,
	}
}

func (v *DBEntryPropertyRichTextValue) ToHTML(w *html.Writer) {
	v.Value.ToHTML(w)
}

type DBEntryPropertyTagsValue struct {
	Values []*DBEntryTag
}

type DBEntryTag struct {
	Name  string
	Color string
}

func NewDBEntryPropertyTagsValue(values []*DBEntryTag) *DBEntryPropertyTagsValue {
	return &DBEntryPropertyTagsValue{
		Values: values,
	}
}

func (v *DBEntryPropertyTagsValue) ToHTML(w *html.Writer) {
	for _, tag := range v.Values {
		if tag.Color != "" && tag.Color != "default" {
			w.PushTag(html.Tag("span").
				Class("badge bg-primary").
				Style(fmt.Sprintf("background-color: %s !important;", tag.Color)))
		} else {
			w.PushTag(html.Tag("span").Class("badge bg-primary"))
		}
		w.WriteLine(html.Escape(tag.Name))
		w.PopTag()
	}
}

type DBEntryPropertyBoolValue struct {
	Value bool
}

func NewDBEntryPropertyBoolValue(value bool) *DBEntryPropertyBoolValue {
	return &DBEntryPropertyBoolValue{
		Value: value,
	}
}

func (v *DBEntryPropertyBoolValue) ToHTML(w *html.Writer) {
	if v.Value {
		w.WriteLine(html.Escape("Yes"))
	} else {
		w.WriteLine(html.Escape("No"))
	}
}
