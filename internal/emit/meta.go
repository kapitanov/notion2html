package emit

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type metadata struct {
	filename string
	json     *metadataJSON
	force    bool
}

type metadataJSON struct {
	Pages map[string]*metadataJSONItem `json:"pages"`
}

type metadataJSONItem struct {
	LastEdited *time.Time `json:"lastEdited"`
}

func (m *metadata) IsUpToDate(pageID string, lastEdited time.Time) bool {
	if m.force {
		return false
	}

	item, ok := m.json.Pages[pageID]
	if !ok {
		return false
	}

	if item.LastEdited == nil {
		return false
	}

	lastKnown := *item.LastEdited

	if lastEdited.Before(lastKnown) || lastEdited.Equal(lastKnown) {
		return true
	}

	return false
}

func (m *metadata) UpdateLastEdited(pageID string, lastEdited time.Time) error {
	item, ok := m.json.Pages[pageID]
	if !ok {
		item = &metadataJSONItem{}
	}

	item.LastEdited = &lastEdited
	m.json.Pages[pageID] = item

	return m.Save()
}

func (m *metadata) Save() error {
	bs, err := json.MarshalIndent(m.json, "", "    ")
	if err != nil {
		return err
	}

	f, err := os.Create(m.filename)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()

	_, err = f.Write(bs)
	if err != nil {
		return err
	}

	return nil
}

func (e *Emitter) loadMetadata(force bool) (*metadata, error) {
	val := &metadataJSON{
		Pages: make(map[string]*metadataJSONItem),
	}

	filename := filepath.Join(e.outputDirectory, "meta.json")
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return nil, err
		}
	} else {
		err = json.Unmarshal(bs, val)
		if err != nil {
			return nil, err
		}
	}

	m := &metadata{
		filename: filename,
		json:     val,
		force:    force,
	}

	return m, nil
}
