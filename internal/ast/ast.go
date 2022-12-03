package ast

import (
	"github.com/kapitanov/notion2html/internal/html"
)

type Node interface {
	ToHTML(w *html.Writer)
}

type Container interface {
	GetNodes() []Node
	AppendNode(node Node)
}

type ContainerEx interface {
	Container
	ShouldProcessChildren() bool
}
