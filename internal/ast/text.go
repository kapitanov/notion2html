package ast

import (
	"fmt"

	"github.com/kapitanov/notion2html/internal/html"
)

type Text struct {
	Runs []*TextRun
}

func NewPlainText(str string) *Text {
	return &Text{
		Runs: []*TextRun{
			{
				Text: str,
			},
		},
	}
}

func NewHrefText(str, url string) *Text {
	return &Text{
		Runs: []*TextRun{
			{
				Text: str,
				Href: url,
			},
		},
	}
}

func (ast *Text) ToHTML(w *html.Writer) {
	for _, el := range ast.Runs {
		el.ToHTML(w)
	}
}

type TextRun struct {
	Bold          bool
	Italic        bool
	Strikethrough bool
	Underline     bool
	Code          bool
	Color         string
	Href          string
	Text          string
}

func (ast *TextRun) ToHTML(w *html.Writer) {
	tagCount := 0

	if ast.Href != "" {
		w.PushTag(html.Tag("a").Href(ast.Href).Inline())
		tagCount++
	}

	if ast.Color != "" && ast.Color != "default" {
		w.PushTag(html.Tag("span").Style(fmt.Sprintf("color: %v;", ast.Color)).Inline())
		tagCount++
	}

	if ast.Bold {
		w.PushTag(html.Tag("strong").Inline())
		tagCount++
	}

	if ast.Italic {
		w.PushTag(html.Tag("em").Inline())
		tagCount++
	}

	if ast.Strikethrough {
		w.PushTag(html.Tag("s").Inline())
		tagCount++
	}

	if ast.Underline {
		w.PushTag(html.Tag("u").Inline())
		tagCount++
	}

	if ast.Code {
		w.PushTag(html.Tag("code").Inline())
		tagCount++
	}

	w.Write(ast.Text)

	for tagCount > 0 {
		w.PopTag()
		tagCount--
	}
}
