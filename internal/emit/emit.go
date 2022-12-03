package emit

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/kapitanov/notion2html/internal/convert"
	"github.com/kapitanov/notion2html/internal/html"
)

type Emitter struct {
	outputDirectory string
	builder         *convert.ASTBuilder
	metadata        *metadata
	pageCount       int
}

func NewEmitter(outputDirectory string, builder *convert.ASTBuilder, force bool) (*Emitter, error) {
	err := os.MkdirAll(outputDirectory, 0777)
	if err != nil && err != os.ErrExist {
		return nil, err
	}

	e := &Emitter{
		outputDirectory: outputDirectory,
		builder:         builder,
	}

	m, err := e.loadMetadata(force)
	if err != nil {
		return nil, err
	}

	e.metadata = m

	return e, nil
}

func (e *Emitter) emit(filename string, fn func(w io.Writer, path string) error) error {
	path := filepath.Join(e.outputDirectory, filename)

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	err = fn(f, path)
	if err != nil {
		return err
	}

	return nil
}

func (e *Emitter) emitHTML(filename string, fn func(w *html.Writer) error) error {
	return e.emit(filename, func(f io.Writer, path string) error {
		directory, _ := filepath.Split(path)

		w := html.NewWriter(f, directory)

		err := fn(w)
		if err != nil {
			return err
		}

		err = w.Complete()
		if err != nil {
			return err
		}

		return nil
	})
}

func (e *Emitter) emitJSON(filename string, v interface{}) error {
	return e.emit(filename, func(w io.Writer, path string) error {
		bs, err := json.MarshalIndent(v, "", "    ")
		if err != nil {
			return err
		}

		_, err = w.Write(bs)
		if err != nil {
			return err
		}

		return nil
	})
}
