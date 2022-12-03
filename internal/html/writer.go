package html

import (
	"bytes"
	pkghtml "html"
	"io"
	"strings"
)

type Writer struct {
	f         io.Writer
	directory string
	err       error
	tags      []T
	indent    int
}

func NewWriter(f io.Writer, directory string) *Writer {
	return &Writer{
		f:         f,
		directory: directory,
	}
}

func (w *Writer) Complete() error {
	return w.err
}

func (w *Writer) WriteLines(strs ...string) {
	for _, str := range strs {
		w.WriteLine(str)
	}
}

func (w *Writer) WriteLine(str string) {
	strs := strings.Split(str, "\n")
	for _, str := range strs {
		for i := 0; i < w.indent; i++ {
			w.Write("\t")
		}

		w.Write(str)
		w.Write("\n")
	}
}

func (w *Writer) WriteLineNoIndent(str string) {
	i := w.indent
	w.indent = 0
	w.WriteLine(str)
	w.indent = i
}

func (w *Writer) Write(str string) {
	if w.err != nil {
		return
	}

	_, w.err = io.WriteString(w.f, str)
	if w.err != nil {
		return
	}
}

func Escape(str string) string {
	return pkghtml.EscapeString(str)
}

func (w *Writer) Render(fn func(w *Writer) error) string {
	if w.err != nil {
		return ""
	}

	buffer := &bytes.Buffer{}
	nestedWriter := &Writer{
		f:         buffer,
		directory: w.directory,
	}

	err := fn(nestedWriter)
	if err != nil {
		w.err = err
		return ""
	}

	str := string(buffer.Bytes())
	return str
}
