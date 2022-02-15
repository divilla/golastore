package e

import "github.com/divilla/golastore/pkg/html/a"

// Main root

func Html(lang string) *E {
	return Elm("html", []a.A{
		{K: "xmlns", V: "http://www.w3.org/1999/xhtml"},
		{K: "xml:lang", V: lang},
		{K: "lang", V: lang},
	})
}

// Document metadata

func Base(attrs ...a.A) *E {
	return Elm("base", attrs)
}

func Head(attrs ...a.A) *E {
	return Elm("head", attrs)
}

func Link(attrs ...a.A) *E {
	return Elm("link", attrs)
}

func Meta(attrs ...a.A) *E {
	return Elm("meta", attrs)
}

func Style(attrs ...a.A) *E {
	return Elm("style", attrs)
}

func Title(attrs ...a.A) *E {
	return Elm("title", attrs)
}

// Sectioning root

func Body(attrs ...a.A) *E {
	return Elm("body", attrs)
}

// Content sectioning

func Address(attrs ...a.A) *E {
	return Elm("address", attrs)
}

func Article(attrs ...a.A) *E {
	return Elm("article", attrs)
}

func Aside(attrs ...a.A) *E {
	return Elm("aside", attrs)
}

func Header(attrs ...a.A) *E {
	return Elm("header", attrs)
}

func Footer(attrs ...a.A) *E {
	return Elm("footer", attrs)
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

func Main(attrs ...a.A) *E {
	return Elm("main", attrs)
}

func Nav(attrs ...a.A) *E {
	return Elm("nav", attrs)
}

func Section(attrs ...a.A) *E {
	return Elm("section", attrs)
}

// Text content

func Blockquote(attrs ...a.A) *E {
	return Elm("blockquote", attrs)
}

func DD(attrs ...a.A) *E {
	return Elm("dd", attrs)
}

func Div(attrs ...a.A) *E {
	return Elm("div", attrs)
}

func DL(attrs ...a.A) *E {
	return Elm("dl", attrs)
}

func DT(attrs ...a.A) *E {
	return Elm("dt", attrs)
}

func Figcaption(attrs ...a.A) *E {
	return Elm("figcaption", attrs)
}

func Figure(attrs ...a.A) *E {
	return Elm("figure", attrs)
}

func HR(attrs ...a.A) *E {
	return Elm("hr", attrs)
}

func Ul(attrs ...a.A) *E {
	return Elm("ul", attrs)
}

func Li(attrs ...a.A) *E {
	return Elm("li", attrs)
}

func Ol(attrs ...a.A) *E {
	return Elm("ol", attrs)
}

func P(attrs ...a.A) *E {
	return Elm("p", attrs)
}

func Pre(attrs ...a.A) *E {
	return Elm("pre", attrs)
}

// Inline text semantics

func A(attrs ...a.A) *E {
	return Elm("a", attrs)
}

func Abbr(attrs ...a.A) *E {
	return Elm("abbr", attrs)
}

func B(attrs ...a.A) *E {
	return Elm("b", attrs)
}

func Bdi(attrs ...a.A) *E {
	return Elm("bdi", attrs)
}

func Bdo(attrs ...a.A) *E {
	return Elm("bdo", attrs)
}

func Br(attrs ...a.A) *E {
	return Elm("br", attrs)
}

func Cite(attrs ...a.A) *E {
	return Elm("cite", attrs)
}

func Code(attrs ...a.A) *E {
	return Elm("code", attrs)
}

func Data(attrs ...a.A) *E {
	return Elm("data", attrs)
}

func Dnf(attrs ...a.A) *E {
	return Elm("dnf", attrs)
}

func Em(attrs ...a.A) *E {
	return Elm("em", attrs)
}

func I(attrs ...a.A) *E {
	return Elm("i", attrs).Open()
}

func Kbd(attrs ...a.A) *E {
	return Elm("kbd", attrs).Open()
}

func Mark(attrs ...a.A) *E {
	return Elm("mark", attrs).Open()
}

func Q(attrs ...a.A) *E {
	return Elm("q", attrs).Open()
}

func S(attrs ...a.A) *E {
	return Elm("s", attrs).Open()
}

func Samp(attrs ...a.A) *E {
	return Elm("samp", attrs).Open()
}

func Small(attrs ...a.A) *E {
	return Elm("small", attrs).Open()
}

func Span(attrs ...a.A) *E {
	return Elm("span", attrs)
}

func Strong(attrs ...a.A) *E {
	return Elm("strong", attrs)
}

func Sub(attrs ...a.A) *E {
	return Elm("sub", attrs)
}

func Sup(attrs ...a.A) *E {
	return Elm("sup", attrs)
}

func Time(attrs ...a.A) *E {
	return Elm("time", attrs)
}

// U underline
func U(attrs ...a.A) *E {
	return Elm("u", attrs)
}

//Image and multimedia

func Area(attrs ...a.A) *E {
	return Elm("area", attrs)
}

func Img(attrs ...a.A) *E {
	return Elm("img", attrs)
}

func Map(attrs ...a.A) *E {
	return Elm("map", attrs)
}

// Scripting

func Script(attrs ...a.A) *E {
	return Elm("script", attrs).Open()
}

// Table content

func Caption(attrs ...a.A) *E {
	return Elm("caption", attrs)
}

func Table(attrs ...a.A) *E {
	return Elm("table", attrs)
}

func Thead(attrs ...a.A) *E {
	return Elm("thead", attrs)
}

func Tbody(attrs ...a.A) *E {
	return Elm("tbody", attrs)
}

func Tfoot(attrs ...a.A) *E {
	return Elm("tfoot", attrs)
}

func Th(attrs ...a.A) *E {
	return Elm("th", attrs)
}

func Td(attrs ...a.A) *E {
	return Elm("td", attrs)
}

func Tr(attrs ...a.A) *E {
	return Elm("tr", attrs)
}

// Forms

func Button(attrs ...a.A) *E {
	return Elm("button", attrs)
}

func Datalist(attrs ...a.A) *E {
	return Elm("datalist", attrs)
}

func Fieldset(attrs ...a.A) *E {
	return Elm("fieldset", attrs)
}

func Form(attrs ...a.A) *E {
	return Elm("form", attrs)
}

func Label(attrs ...a.A) *E {
	return Elm("label", attrs)
}

func Input(attrs ...a.A) *E {
	return Elm("input", attrs)
}

func Textarea(attrs ...a.A) *E {
	return Elm("textarea", attrs)
}

func Select(attrs ...a.A) *E {
	return Elm("select", attrs)
}

func Option(attrs ...a.A) *E {
	return Elm("option", attrs)
}
