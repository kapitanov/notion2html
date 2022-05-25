package html

import (
	"bytes"
	"html/template"
	"io"
	"log"
	"path/filepath"
	"time"
)

func (w *Writer) Template(name string, model interface{}) {
	w.renderTemplateTo(w.f, name, model)
}

func (w *Writer) RenderTemplate(name string, model interface{}) string {
	buffer := &bytes.Buffer{}
	w.renderTemplateTo(buffer, name, model)
	if w.err != nil {
		return ""
	}

	str := string(buffer.Bytes())
	return str
}

func (w *Writer) renderTemplateTo(wr io.Writer, name string, model interface{}) {
	if w.err != nil {
		return
	}

	filename := filepath.Join("./templates", name)

	t := template.New(name)
	t = t.Funcs(templateFunctions)
	t, w.err = t.ParseFiles(filename)
	if w.err != nil {
		log.Printf("unable to parse template \"%s\": %s", filename, w.err)
		return
	}

	w.err = t.Execute(wr, model)
	if w.err != nil {
		log.Printf("unable to execute template \"%s\": %s", filename, w.err)
		return
	}
}

var templateFunctions = template.FuncMap{
	"raw": func(s string) (template.HTML, error) {
		return template.HTML(s), nil
	},
	"formatTime": func(t time.Time, format string, args ...interface{}) string {
		return t.Format(format)
	},
	"mul": func(x, y int) int {
		return x*y
	},
}
