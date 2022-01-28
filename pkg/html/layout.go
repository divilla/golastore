package html

import (
	"strings"
)

type (
	Renderer interface {
		Render(depth int, bb *strings.Builder)
	}

	Layout struct {
		children []Renderer
		view     IView
	}

	ILayout interface {
		View(v IView) *Layout
		Render(depth int, bb *strings.Builder)
	}
)

func NewLayout(children ...Renderer) *Layout {
	return &Layout{
		children: children,
	}
}

func (l *Layout) View(v IView) *Layout {
	l.view = v
	return l
}

func (l *Layout) Render(depth int, bb *strings.Builder) {
	for k, elm := range l.children {
		if k > 0 {
			bb.WriteString("\n")
		}
		elm.Render(depth, bb)
	}
}
