package ast

import (
	"strings"

	"github.com/kapitanov/notion2html/internal/html"
)

type Video struct {
	URL         string
	Caption     *Text
	CaptionHTML string
	Youtube     *YoutubeVideo
}

type YoutubeVideo struct {
	URL string
}

func (ast *Video) ToHTML(w *html.Writer) {
	ast.CaptionHTML = w.Render(func(wr *html.Writer) error {
		ast.Caption.ToHTML(wr)
		return nil
	})

	youtubeURL := ""
	if strings.HasPrefix(ast.URL, "https://www.youtube.com/") {
		youtubeURL = strings.TrimPrefix(ast.URL, "https://www.youtube.com/")
	} else if strings.HasPrefix(ast.URL, "https://youtube.com/") {
		youtubeURL = strings.TrimPrefix(ast.URL, "https://youtube.com/")
	} else if strings.HasPrefix(ast.URL, "https://youtu.be/") {
		youtubeURL = strings.TrimPrefix(ast.URL, "https://youtu.be/")
	}
	if youtubeURL != "" {
		ast.Youtube = &YoutubeVideo{
			URL: "https://www.youtube.com/embed/" + youtubeURL,
		}
	}

	w.Template("video.html", ast)
}
