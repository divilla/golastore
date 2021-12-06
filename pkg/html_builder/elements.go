package html_builder

func (d *D) Inner(inner string) E {
	return d.Elm("", nil, inner)
}

func (d *D) Html(lang string, children ...E) E {
	return d.Elm("html", A{
		"xmlns": "http://www.w3.org/1999/xhtml",
		"xml:lang": lang,
		"lang": lang,
	}, "", children...)
}

func (d *D) Head(children ...E) E {
	return d.Elm("head", nil, "", children...)
}

func (d *D) Title(inner string) E {
	return d.Elm("title", nil, inner)
}

func (d *D) Meta(attrs A) E {
	return d.Elm("meta", attrs, "")
}

func (d *D) Link(attrs A) E {
	return d.Elm("link", attrs, "")
}

func (d *D) Script(attrs A, inner ...string) E {
	i := ""
	if len(inner) > 0 {
		i = inner[0]
	}

	e := d.Elm("script", attrs, i)
	e.open = true
	return e
}

func (d *D) Body(children ...E) E {
	return d.Elm("body", nil, "", children...)
}

func (d *D) Div(class string, children ...E) E {
	return d.Elm("div", A{}.Class(class), "", children...)
}

func (d *D) Table(class string, children ...E) E {
	return d.Elm("table", A{}.Class(class), "", children...)
}

func (d *D) Thead(children ...E) E {
	return d.Elm("thead", nil, "", children...)
}

func (d *D) Tbody(children ...E) E {
	return d.Elm("tbody",nil,"", children...)
}

func (d *D) Tr(children ...E) E {
	return d.Elm("tr",nil,"", children...)
}

func (d *D) Th(inner string, attrs ...A) E {
	var a A
	if len(attrs) > 0 {
		a = attrs[0]
	}
	return d.Elm("th", a, inner)
}

func (d *D) Td(inner string, attrs ...A) E {
	var a A
	if len(attrs) > 0 {
		a = attrs[0]
	}
	return d.Elm("td", a, inner)
}

func (d *D) Caption(inner string, attrs ...A) E {
	var a A
	if len(attrs) > 0 {
		a = attrs[0]
	}
	return d.Elm("caption", a, inner)
}

func (d *D) H1(inner string, attrs ...A) E {
	var a A
	if len(attrs) > 0 {
		a = attrs[0]
	}
	return d.Elm("h1", a, inner)
}

func (d *D) H2(inner string, attrs ...A) E {
	var a A
	if len(attrs) > 0 {
		a = attrs[0]
	}
	return d.Elm("h2", a, inner)
}

func (d *D) H3(inner string, attrs ...A) E {
	var a A
	if len(attrs) > 0 {
		a = attrs[0]
	}
	return d.Elm("h3", a, inner)
}

func (d *D) H4(inner string, attrs ...A) E {
	var a A
	if len(attrs) > 0 {
		a = attrs[0]
	}
	return d.Elm("h3", a, inner)
}
