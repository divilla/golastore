package e

import (
	"bytes"
	"github.com/divilla/golastore/pkg/html/a"
)

type (
	E struct {
		Tag      string
		Attrs    []a.A
		Text     string
		Children []*E
		Hidden   bool
		open     bool
	}
)


func Elm(tag string, attrs []a.A) *E {
	return &E{
		Tag: tag,
		Attrs: attrs,
	}
}

func (e *E) T(text string) *E {
	e.Text = text
	return e
}

func (e *E) C(children ...*E) *E {
	e.Children = children
	return e
}

func (e *E) Open() *E {
	e.open = true
	return e
}

func (e E) Render(depth int, bb *bytes.Buffer) {
	if e.Hidden {
		return
	}

	tabs(depth, bb)
	bb.WriteString(`<` + e.Tag)

	for _, attr := range e.Attrs {
		attr.Render(bb)
	}

	if !e.open && e.Text == "" && len(e.Children) == 0 {
		bb.WriteString(" />")
		return
	}
	bb.WriteString(">" + e.Text)

	if len(e.Children) > 0 {
		bb.WriteString("\n")
		for _, elm := range e.Children {
			elm.Render(depth + 1, bb)
			bb.WriteString("\n")
		}
		tabs(depth, bb)
	}

	bb.WriteString("</" + e.Tag + ">")
}

func tabs(depth int, bb *bytes.Buffer) {
	for i := 0; i < depth; i++ {
		bb.WriteString("    ")
	}
}