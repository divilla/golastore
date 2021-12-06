package html

import (
	"bytes"
	"github.com/divilla/golastore/pkg/html/e"
)

type (
	D struct {
		children []*e.E
	}
)

func NewDocument(children ...*e.E) *D {
	return &D{
		children: children,
	}
}

func (d *D) Render() []byte {
	bb := bytes.NewBufferString("<!DOCTYPE html>\n")

	for _, elm := range d.children {
		elm.Render(0, bb)
		bb.WriteString("\n")
	}

	return bb.Bytes()
}
