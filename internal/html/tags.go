package html

import (
	"fmt"
	"strings"
)

type T struct {
	name   string
	attrs  []string
	inline bool
}

func Tag(name string) T {
	return T{
		name: name,
	}
}

func Tagf(format string, args ...interface{}) T {
	name := fmt.Sprintf(format, args...)
	return Tag(name)
}

func (tag T) Attr(key, value string) T {
	if key != "" && value != "" {
		tag.attrs = append(tag.attrs, fmt.Sprintf("%s=\"%s\"", key, Escape(value)))
	}
	return tag
}

func (tag T) ID(value string) T {
	return tag.Attr("id", value)
}

func (tag T) Class(value string) T {
	return tag.Attr("class", value)
}

func (tag T) Style(value string) T {
	return tag.Attr("style", value)
}

func (tag T) Href(value string) T {
	return tag.Attr("href", value)
}

func (tag T) Inline() T {
	tag.inline = true
	return tag
}

func (tag T) writeStart(w *Writer) {
	str := "<" + tag.name
	if len(tag.attrs) > 0 {
		str += " " + strings.Join(tag.attrs, " ")
	}
	str += ">"

	if tag.inline {
		w.Write(str)
	} else {
		w.indent++
		w.WriteLine(str)
	}
}

func (tag T) writeEnd(w *Writer) {
	str := fmt.Sprintf("</%s>", tag.name)

	if tag.inline {
		w.Write(str)
	} else {
		w.WriteLine(str)
		w.indent--
	}
}

func (w *Writer) PushTag(tag T) {
	tag.writeStart(w)
	w.tags = append(w.tags, tag)
}

func (w *Writer) PopTag() {
	if len(w.tags) > 0 {
		tag := w.tags[len(w.tags)-1]
		w.tags = w.tags[:len(w.tags)-1]

		tag.writeEnd(w)
	}
}

func (w *Writer) WithTag(tag T, fn func() error) {
	if w.err != nil {
		return
	}

	w.PushTag(tag)
	w.err = fn()
	w.PopTag()
}
