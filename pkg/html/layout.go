package html

import (
	"github.com/tidwall/gjson"
	"strings"
)

type (
	Renderer interface {
		Render(depth int, bb *strings.Builder)
	}

	Layout struct {
		content  Renderer
		children []Renderer
		data     []byte
	}
)

func NewLayout(children ...Renderer) *Layout {
	return &Layout{
		children: children,
	}
}

func (l *Layout) Content(c Renderer) *Layout {
	l.content = c
	return l
}

func (l *Layout) Data(d []byte) *Layout {
	l.data = d
	return l
}

func (l *Layout) Get(path string) string {
	return gjson.GetBytes(l.data, path).String()
}

func (l *Layout) Render(depth int, bb *strings.Builder) {
	for k, elm := range l.children {
		if k > 0 {
			bb.WriteString("\n")
		}
		elm.Render(depth, bb)
	}
}
