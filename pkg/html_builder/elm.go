package html_builder

import (
	"bytes"
)

type (
	E struct {
		Tag      string
		Attrs    A
		Show     bool
		Inner    string
		Children *[]E
		open     bool
		buffer   *bytes.Buffer
		depth    int
	}

	A map[string]string
)

func (d *D) Elm(tag string, attributes A, innerHTML string, children ...E) E {
	return E{
		Tag:      tag,
		Attrs:    attributes,
		Show:     true,
		Inner:    innerHTML,
		Children: &children,
		buffer:   d.buffer,
	}
}

//InnerHTML sets inner html
func (e E) InnerHTML(val string, show ...bool) E {
	if len(show) > 0 && show[0] == false {
		return e
	}

	e.Inner = val

	return e
}

func (e E) If(show bool) E {
	e.Show = show
	return e
}

func (e E) Class(class string, show ...bool) E {
	if len(show) > 0 && show[0] == false {
		return e
	}

	if e.Attrs == nil {
		e.Attrs = A{"class": class}
	} else if _, ok := e.Attrs["class"]; !ok {
		e.Attrs["class"] = class
	} else {
		e.Attrs["class"] += " " + class
	}

	return e
}

func (e E) Attr(key string, val string, show ...bool) E {
	if len(show) > 0 && show[0] == false {
		return e
	}

	if e.Attrs == nil {
		e.Attrs = A{key: val}
	} else {
		e.Attrs[key] = val
	}

	return e
}

func (e E) Child(child E, show ...bool) E {
	if len(show) > 0 && show[0] == false {
		return e
	}

	*e.Children = append(*e.Children, child)

	return e
}

func (e E) Bytes() {
	if !e.Show {
		return
	}

	if e.Tag == "" {
		e.buffer.WriteString(e.Inner)
		return
	}

	for i := 0; i < e.depth; i++ {
		e.buffer.WriteString("\t")
	}

	e.buffer.WriteString(`<` + e.Tag)

	for k, v := range e.Attrs {
		e.buffer.WriteString(" " + k + `="` + v + `"`)
	}

	if !e.open && e.Inner == "" && len(*e.Children) == 0 {
		e.buffer.WriteString(" />")
		return
	}

	e.buffer.WriteString(">" + e.Inner)
	if len(*e.Children) > 0 {
		e.buffer.WriteString("\n")
		for _, elm := range *e.Children {
			elm.depth = e.depth + 1
			elm.Bytes()
			e.buffer.WriteString("\n")
		}

		for i := 0; i < e.depth; i++ {
			e.buffer.WriteString("\t")
		}
	}

	e.buffer.WriteString("</" + e.Tag + ">")
}

//Class sets class attribute
func (a A) Class(class string, apply ...bool) A {
	if len(apply) > 0 && apply[0] == false {
		return a
	}

	if _, ok := a["class"]; !ok {
		a["class"] = class
	} else {
		a["class"] += " " + class
	}

	return a
}

func (a A) Attr(key string, val string, apply ...bool) A {
	if len(apply) > 0 && apply[0] == false {
		return a
	}

	a[key] = val

	return a
}
