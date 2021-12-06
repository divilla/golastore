package e

import "github.com/divilla/render/pkg/html/a"

func Html(lang string) *E {
	return Elm("html", []a.A{
		{K: "xmlns", V: "http://www.w3.org/1999/xhtml"},
		{K: "xml:lang", V: lang},
		{K: "lang", V: lang},
	})
}

func Head(attrs ...a.A) *E {
	return Elm("head", attrs)
}

func Title(attrs ...a.A) *E {
	return Elm("title", attrs)
}

func Meta(attrs ...a.A) *E {
	return Elm("meta", attrs)
}

func Link(attrs ...a.A) *E {
	return Elm("link", attrs)
}

func Script(attrs ...a.A) *E {
	return Elm("script", attrs).Open()
}

func Body(attrs ...a.A) *E {
	return Elm("body", attrs)
}

func Section(attrs ...a.A) *E {
	return Elm("section", attrs)
}

func Strong(attrs ...a.A) *E {
	return Elm("strong", attrs)
}

func P(attrs ...a.A) *E {
	return Elm("p", attrs)
}

func Div(attrs ...a.A) *E {
	return Elm("div", attrs)
}

func Table(attrs ...a.A) *E {
	return Elm("table", attrs)
}

func Thead(attrs ...a.A) *E {
	return Elm("thead", attrs)
}

func Tbody(attrs ...a.A) *E {
	return Elm("tbody",attrs)
}

func Tr(attrs ...a.A) *E {
	return Elm("tr",attrs)
}

func Th(attrs ...a.A) *E {
	return Elm("th", attrs)
}

func Td(attrs ...a.A) *E {
	return Elm("td", attrs)
}

func Caption(attrs ...a.A) *E {
	return Elm("caption", attrs)
}

func H1(attrs ...a.A) *E {
	return Elm("h1", attrs)
}

func H2(attrs ...a.A) *E {
	return Elm("h2", attrs)
}

func H3(attrs ...a.A) *E {
	return Elm("h3", attrs)
}

func H4(attrs ...a.A) *E {
	return Elm("h4", attrs)
}

func H5(attrs ...a.A) *E {
	return Elm("h5", attrs)
}

func H6(attrs ...a.A) *E {
	return Elm("h6", attrs)
}
