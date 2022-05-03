package components

import (
	"github.com/divilla/golastore/internal/domain_model"
	"github.com/divilla/golastore/pkg/html"
	"github.com/divilla/golastore/pkg/html/a"
	"github.com/divilla/golastore/pkg/html/e"
)

func NewProductList(items []*domain_model.ProductListItem) []html.Renderer {
	elms := make([]html.Renderer, len(items))
	for key, item := range items {
		elms[key] = NewProductListItem(item)
	}
	return elms
}

func NewProductListItem(item *domain_model.ProductListItem) *e.E {
	return e.Div(a.Class("p-2 cols-xl-5 cols-lg-4 cols-md-3 cols-sm-2 cols-1 card")).Children(
		e.Div(a.Class("p-2 is-flex is-justify-content-space-between")).Children(
			e.Span(a.Class("tag is-primary is-medium strong")).
				Text("Promo"),
			e.Span(a.Class("tag is-info is-medium strong")).
				Text("Best Buy"),
			e.Span(a.Class("tag is-danger is-medium strong")).
				Text("Hot"),
		),
		e.Div(a.Class("card-image")).Children(
			e.A(a.Href("#"), a.Class("list-image")).Children(
				e.Img(a.Src(item.ImageURL()), a.Title(item.Name), a.Alt(item.Name)),
			),
		),
		e.Div(a.Class("p-2 is-flex is-justify-content-space-between")).Children(
			e.Span(a.Class("line-through")).
				Text(item.OldPriceFormat()),
			e.Span(a.Class("tag is-warning is-medium strong")).
				Text(item.Discount()),
			e.Span(a.Class("strong")).
				Text(item.PriceFormat()),
		),
		e.Div(a.Class("p-2 is-flex is-justify-content-space-between")).Children(
			e.Button(a.Class("button is-small is-info is-light")).Children(
				e.Span(a.Class("icon is-small")).Children(
					e.I(a.Class("fas fa-shopping-cart is-danger")),
				),
			),
			e.Div(a.Class("field has-addons")).Children(
				e.P(a.Class("control")).Children(
					e.Button(a.Class("button is-small")).Children(
						e.Span(a.Class("icon is-small")).Children(
							e.I(a.Class("fas fa-minus")),
						),
					),
				),
				e.P(a.Class("control")).Children(
					e.Button(a.Class("button is-small")).Children(
						e.Span().Text("0"),
					),
				),
				e.P(a.Class("control")).Children(
					e.Button(a.Class("button is-small")).Children(
						e.Span(a.Class("icon is-small")).Children(
							e.I(a.Class("fas fa-plus")),
						),
					),
				),
			),
			e.Button(a.Class("button is-small is-danger is-light")).Children(
				e.Span(a.Class("icon is-small")).Children(
					e.I(a.Class("far fa-heart is-danger")),
				),
			),
		),
		e.A(a.Href(item.LinkToProduct()), a.Class("m-2 three-lines-text strong")).
			Text(item.Name),
	)
}
