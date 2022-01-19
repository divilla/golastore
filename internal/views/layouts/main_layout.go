package layouts

import (
	"github.com/divilla/golastore/pkg/html"
	"github.com/divilla/golastore/pkg/html/a"
	"github.com/divilla/golastore/pkg/html/d"
	"github.com/divilla/golastore/pkg/html/e"
	"github.com/tidwall/gjson"
	"strings"
)

func MainLayout(sb *strings.Builder, jr *gjson.Result) {
	html.NewLayout(
		d.Block("<!DOCTYPE html>"),
		e.Html("en-US").Children(
			e.Head().Children(
				e.Meta(a.A{K: "charset", V: "utf-8"}),
				e.Meta(a.A{K: "name", V: "viewport"}, a.A{K: "content", V: "width=device-width, initial-scale=1"}),
				e.Title().Text(jr.Get("mod.title").String()),
				e.Link(a.A{K: "shortcut icon", V: "https://fdn.gsmarena.com/imgroot/static/favicon.ico"}),
				e.Link(a.Rel("stylesheet"), a.Href("/assets/css/style.css")),
			),
			e.Body().Children(
				e.Section(a.Class("section"), a.Style("padding-top: 1.5rem")).Children(
					e.Div(a.Class("columns")).Children(
						e.Div(a.Class("column is-2")).Children(
							e.Img(a.Src("/assets/images/market-logo.png"), a.Width("142"), a.Height("42")),
						),
						e.Div(a.Class("column is-5")).Children(
							e.Div(a.Class("field")).Children(
								e.P(a.Class("control has-icons-right")).Children(
									e.Input(a.Class("input is-rounded"), a.Type("text"), a.Placeholder("Search")),
									e.Span(a.Class("icon is-small is-right has-text-danger-dark")).Children(
										e.I(a.Class("fas fa-search")),
									),
								),
							),
						),
						e.Div(a.Class("column is-5 is-flex is-justify-content-flex-end")).Children(
							e.P(a.Class("buttons")).Children(
								e.A(a.Class("button is-white")).Children(
									e.Span(a.Class("icon is-small has-text-success-dark")).Children(
										e.I(a.Class("fas fa-sign-in-alt")),
									),
									e.Span().Text("Sign in"),
								),
								e.A(a.Class("button is-white")).Children(
									e.Span(a.Class("icon is-small has-text-info-dark")).Children(
										e.I(a.Class("fas fa-user-plus")),
									),
									e.Span().Text("Sign up"),
								),
								e.A(a.Class("button is-white")).Children(
									e.Span(a.Class("icon is-small has-text-danger-dark")).Children(
										e.I(a.Class("fas fa-sign-out-alt")),
									),
									e.Span().Text("Sign out"),
								),
								e.A(a.Class("button is-dark")).Children(
									e.Span(a.Class("badge is-danger"), a.Title("Cart Items")).Text("8"),
									e.Span().Text("0,00 kn"),
									e.Span(a.Class("icon is-small has-text-warning")).Children(
										e.I(a.Class("fas fa-shopping-cart")),
									),
								),
							),
						),
					),
					//e.Nav(a.Class("navbar"), a.Role("navigation"), a.AriaLabel("main-navigation")).Children(
					//	e.Div(a.Class("navbar-brand")).Children(
					//		e.A(a.Class("navbar-item"), a.Href("/")).Children(
					//			e.Img(a.Src("/assets/images/market-logo.png"), a.Width("142"), a.Height("42")),
					//		),
					//		e.A(a.Class("navbar-burger"), a.Role("button"), a.AriaExpanded("false"), a.DataTarget("navbarBasicExample")).Children(
					//			e.Span(a.AriaHidden("true")),
					//			e.Span(a.AriaHidden("true")),
					//			e.Span(a.AriaHidden("true")),
					//		),

					//	),
					//	e.Div(a.Id("navbarBasicExample"), a.Class("navbar-menu")).Children(
					//		e.Div(a.Class("navbar-start")).Children(
					//			e.A(a.Class("navbar-item")).T(
					//				"Home",
					//			),
					//			e.A(a.Class("navbar-item")).T(
					//				"Documentation",
					//			),
					//		),
					//	),
					//),
					//container(nil)...
				),
				e.Section(a.Class("section"), a.Style("padding-top: 1.5rem")).Children(
					d.If(false,
						d.Block("This is text"),
					).Else(
						d.Block(`<div style="color:red">This is text</div>`),
					),
				),
			),
		),
	).Render(0, sb)
}
