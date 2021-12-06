package a

import "bytes"

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

// Render converts attribute to string.
func (a A) Render(bb *bytes.Buffer) {
	bb.WriteString(" " + a.K + "=\"" + a.V + "\"")
}

func Class(val string) A {
	return Attr("class", val)
}

func Href(val string) A {
	return Attr("href", val)
}

func Src(val string) A {
	return Attr("src", val)
}

func Rel(val string) A {
	return Attr("rel", val)
}

func Integrity(val string) A {
	return Attr("integrity", val)
}

func Crossorigin(val string) A {
	return Attr("crossorigin", val)
}
