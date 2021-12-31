package e

import (
	"bytes"
	"github.com/divilla/golastore/pkg/html"
	"github.com/divilla/golastore/pkg/html/a"
)

type (
	E struct {
		Tag      string
		Attrs    []a.A
		Text     string
		Hidden   bool
		children []html.Renderer
		open     bool
	}
)

func Elm(tag string, attrs []a.A) *E {
	return &E{
		Tag:   tag,
		Attrs: attrs,
	}
}

func (e *E) T(text string) *E {
	e.Text = text
	return e
}

func (e *E) Children(children ...html.Renderer) *E {
	e.children = children
	return e
}

func (e *E) Open() *E {
	e.open = true
	return e
}

func (e *E) Render(depth int, bb *bytes.Buffer) {
	if e.Hidden {
		return
	}

	html.Tabs(depth, bb)
	bb.WriteString(`<` + e.Tag)

	for _, attr := range e.Attrs {
		attr.Render(bb)
	}

	if !e.open && e.Text == "" && len(e.children) == 0 {
		bb.WriteString(" />")
		return
	}
	bb.WriteString(">" + e.Text)

	if len(e.children) > 0 {
		bb.WriteString("\n")
		for _, elm := range e.children {
			elm.Render(depth+1, bb)
			bb.WriteString("\n")
		}
		html.Tabs(depth, bb)
	}

	bb.WriteString("</" + e.Tag + ">")
}
