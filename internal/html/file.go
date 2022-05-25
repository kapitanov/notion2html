package html

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"hash"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func (w *Writer) CacheURL(href string) string {
	if w.err != nil {
		return ""
	}

	filename, err := w.downloadFile(href)
	if err != nil {
		w.err = err
		return ""
	}

	return filename
}

func (w *Writer) downloadFile(href string) (string, error) {
	if w.err != nil {
		return "", w.err
	}

	resp, err := http.Get(href)
	if err != nil {
		w.err = err
		return "", err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, resp.Body)
	if err != nil {
		w.err = err
		return "", err
	}

	h := sha1.New()
	_, err = h.Write(buffer.Bytes())
	if err != nil {
		w.err = err
		return "", err
	}

	filename := w.generateCachedFilename(h, href)
	log.Printf("caching %s", filename)
	f, err := os.Create(filepath.Join(w.directory, filename))
	if err != nil {
		w.err = err
		return "", err
	}
	defer func() {
		_ = f.Close()
	}()

	_, err = io.Copy(f, &buffer)
	if err != nil {
		w.err = err
		return "", err
	}

	return filename, nil
}

func (w *Writer) generateCachedFilename(h hash.Hash, href string) string {
	parts := strings.SplitN(href, "?", 2)
	href = parts[0]

	filename := hex.EncodeToString(h.Sum(nil))
	filename = filename + filepath.Ext(href)

	return filename
}