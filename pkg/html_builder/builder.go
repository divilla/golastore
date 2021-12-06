package html_builder

import (
	"bytes"
)

type D struct {
	Children []E
	buffer   *bytes.Buffer
}

func Build(elements ...E) *D {
	return &D{
		Children: elements,
		buffer: new(bytes.Buffer),
	}
}

func (d *D) Child(element E) {
	d.Children = append(d.Children, element)
}

func (d *D) Bytes() []byte {
	d.buffer.WriteString("<!DOCTYPE html>\n")
	for _, elm := range d.Children {
		elm.Bytes()
	}

	return d.buffer.Bytes()
}
