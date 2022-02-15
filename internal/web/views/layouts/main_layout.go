package layouts

import (
	"github.com/divilla/golastore/internal/view_model"
	"github.com/divilla/golastore/pkg/html"
	"github.com/divilla/golastore/pkg/html/a"
	"github.com/divilla/golastore/pkg/html/d"
	"github.com/divilla/golastore/pkg/html/e"
	"github.com/divilla/golastore/pkg/random"
)

type (
	IMainLayoutData interface {
		WebPage() *view_model.WebPage
	}
)

func NewMainLayout(data IMainLayoutData, view html.IView) *html.Layout {
	rnd := random.MustString(32)

	return html.NewLayout(
		d.Block("<!DOCTYPE html>"),
		e.Html("en-US").Children(
			e.Head().Children(
				e.Meta(a.A{K: "charset", V: "utf-8"}),
				e.Meta(a.A{K: "name", V: "viewport"}, a.A{K: "content", V: "width=device-width, initial-scale=1"}),
				e.Title().Text(data.WebPage().MetaTitle),
				e.Link(a.Rel("stylesheet"), a.Href("https://cdnjs.cloudflare.com/ajax/libs/flexboxgrid/6.3.1/flexboxgrid.min.css"), a.Integrity("ssha512-YHuwZabI2zi0k7c9vtg8dK/63QB0hLvD4thw44dFo/TfBFVVQOqEG9WpviaEpbyvgOIYLXF1n7xDUfU3GDs0sw=="), a.CrossOrigin("anonymous"), a.RefererPolicy("no-referrer")),
				e.Link(a.Rel("stylesheet"), a.Href("https://cdnjs.cloudflare.com/ajax/libs/MaterialDesign-Webfont/6.5.95/css/materialdesignicons.css"), a.Integrity("sha512-x6bCTnc3rzkrdNmyFuJBWwOF23tzxLalBSQ2N3BEYKLWOpjfpYpnBrE0wlRHI/j4Oyv72/VuAewe2DmEl/NXdA=="), a.CrossOrigin("anonymous"), a.RefererPolicy("no-referrer")),
				e.Link(a.Rel("stylesheet"), a.Href("/assets/css/style.css?id="+rnd)),
			),
			e.Body().Children(
				e.Section(a.Class("section"), a.Style("padding: 1.5rem")).Children(
					e.Div(a.Class("columns")).Children(
						e.Div(a.Class("column is-2")).Children(
							e.Img(a.Src("/assets/images/market-logo.png"), a.Width("142"), a.Height("42")),
						),
						e.Div(a.Class("column is-4")).Children(
							e.Div(a.Class("field")).Children(
								e.P(a.Class("control has-icons-right")).Children(
									e.Input(a.Class("input is-rounded"), a.Type("text"), a.Placeholder("Search")),
									e.Span(a.Class("icon is-small is-right has-text-danger-dark")).Children(
										e.I(a.Class("fas fa-search")),
									),
								),
							),
						),
						e.Div(a.Class("column is-6 is-flex is-justify-content-flex-end")).Children(
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
				),
				e.HR(a.Style("margin: 0 1.5rem;")),
				e.Section(a.Class("section"), a.Style("padding: 1.5rem;")).Children(
					view,
				),
			),
		),
	)
}
