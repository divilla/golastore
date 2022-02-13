package a

import (
	"strings"
)

// A is the attribute definition type
type (
	A struct {
		K string
		V string
	}
)

// Attr creates new attribute.
func Attr(key string, val string) A {
	return A{
		K: key,
		V: val,
	}
}

// Render writes attribute to a buffer
func (a *A) Render(bb *strings.Builder) {
	if a.K == "" {
		return
	}

	bb.WriteString(" ")
	bb.WriteString(a.K)
	bb.WriteString(`="`)
	bb.WriteString(a.V)
	bb.WriteString(`"`)
}
