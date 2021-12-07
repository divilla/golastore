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

func AriaExpanded(val string) A {
	return Attr("aria-expanded", val)
}

func AriaHidden(val string) A {
	return Attr("aria-hidden", val)
}

func AriaLabel(val string) A {
	return Attr("aria-label", val)
}

func Class(val string) A {
	return Attr("class", val)
}

func Crossorigin(val string) A {
	return Attr("crossorigin", val)
}

func DataTarget(val string) A {
	return Attr("data-target", val)
}

func Height(val string) A {
	return Attr("height", val)
}

func Href(val string) A {
	return Attr("href", val)
}

func Id(val string) A {
	return Attr("id", val)
}

func Integrity(val string) A {
	return Attr("integrity", val)
}

func Placeholder(val string) A {
	return Attr("placeholder", val)
}

func Rel(val string) A {
	return Attr("rel", val)
}

func Role(val string) A {
	return Attr("role", val)
}

func Src(val string) A {
	return Attr("src", val)
}

func Style(val string) A {
	return Attr("style", val)
}

func Title(val string) A {
	return Attr("title", val)
}

func Type(val string) A {
	return Attr("type", val)
}

func Width(val string) A {
	return Attr("width", val)
}
