package html

import (
	"bytes"
)

type (
	Renderer interface {
		Render(depth int, bb *bytes.Buffer)
	}

	D struct {
		children []Renderer
	}
)

func NewDocument(children ...Renderer) *D {
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
