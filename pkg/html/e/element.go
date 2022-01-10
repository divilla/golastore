package e

import (
	"github.com/divilla/golastore/pkg/html"
	"github.com/divilla/golastore/pkg/html/a"
	"strings"
)

type (
	E struct {
		Tag      string
		Attrs    []a.A
		Hidden   bool
		text     string
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

func (e *E) Text(text string) html.Renderer {
	e.text = text
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

func (e *E) Render(depth int, bb *strings.Builder) {
	if e.Hidden {
		return
	}

	html.Tabs(depth, bb)
	bb.WriteString(`<`)
	bb.WriteString(e.Tag)
	for _, attr := range e.Attrs {
		attr.Render(bb)
	}

	if len(e.text) > 0 {
		bb.WriteString(">")
		bb.WriteString(e.text)
		bb.WriteString("</")
		bb.WriteString(e.Tag)
		bb.WriteString(">")
		return
	}

	if len(e.children) > 0 {
		bb.WriteString(`>`)
		for _, elm := range e.children {
			bb.WriteString("\n")
			elm.Render(depth+1, bb)
		}
		bb.WriteString("\n")
		html.Tabs(depth, bb)
		bb.WriteString("</")
		bb.WriteString(e.Tag)
		bb.WriteString(">")
		return
	}

	if e.open {
		bb.WriteString(">")
		bb.WriteString("</")
		bb.WriteString(e.Tag)
		bb.WriteString(">")
		return
	} else {
		bb.WriteString(" />")
		return
	}

}
