package catalog

import (
	"github.com/divilla/golastore/internal/web/catalog/catalog_service"
	"github.com/divilla/golastore/internal/web/views/components/pagination"
	"github.com/divilla/golastore/internal/web/views/layouts"
	"github.com/divilla/golastore/pkg/html"
	"github.com/divilla/golastore/pkg/html/a"
	"github.com/divilla/golastore/pkg/html/e"
)

func NewCategoryView(m *catalog_service.CatalogCategoryModel) *html.Layout {
	paginationComponent := pagination.New(m)
	items := make([]html.Renderer, len(m.ListProducts()))
	for key, item := range m.ListProducts() {
		items[key] =
			e.Div(a.Class("column is-2")).Children(
				e.Div(a.Class("card")).Children(
					e.Div(a.Class("card-content"), a.Style("padding: 0.5rem 0.5rem 0 0.5rem")).Children(
						e.Div(a.Class("content")).Children(
							e.Table(a.Class("table-center-middle")).Children(
								e.Tr().Children(
									e.Td().Children(
										e.Span(a.Class("tag is-primary is-medium strong")).
											Text("Promo"),
									),
									e.Td().Children(
										e.Span(a.Class("tag is-info is-medium strong")).
											Text("Best Buy"),
									),
									e.Td().Children(
										e.Span(a.Class("tag is-danger is-medium strong")).
											Text("Hot"),
									),
								),
							),
						),
					),
					e.Div(a.Class("card-image")).Children(
						e.A(a.Href("#"), a.Class("list-image")).Children(
							e.Img(a.Src(item.ImageURL()), a.Title(item.Name), a.Alt(item.Name)),
						),
					),
					e.Div(a.Class("card-content")).Children(
						e.Div(a.Class("content")).Children(
							e.Table(a.Class("table-center-middle")).Children(
								e.Tr().Children(
									e.Td().Children(
										e.Span(a.Class("line-through")).
											Text(item.OldPriceFormat()),
									),
									e.Td().Children(
										e.Span(a.Class("tag is-warning is-medium strong")).
											Text(item.Discount()),
									),
									e.Td().Children(
										e.Span(a.Class("strong")).
											Text(item.PriceFormat()),
									),
								),
							),
							e.Table(a.Class("table-center")).Children(
								e.Tr().Children(
									e.Td(a.Class("text-align: left")).Children(
										e.P(a.Class("buttons")).Children(
											e.Button(a.Class("button is-small is-info is-light")).Children(
												e.Span(a.Class("icon is-small")).Children(
													e.I(a.Class("fas fa-shopping-cart is-danger")),
												),
											),
										),
									),
									e.Td().Children(
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
									),
									e.Td(a.Class("text-align: right")).Children(
										e.P(a.Class("buttons")).Children(
											e.Button(a.Class("button is-small is-danger is-light")).Children(
												e.Span(a.Class("icon is-small")).Children(
													e.I(a.Class("far fa-heart is-danger")),
												),
											),
										),
									),
								),
							),
							e.A(a.Href(item.LinkToProduct()), a.Class("two-lines-text strong")).
								Text(item.Name),
						),
					),
				),
			)
	}

	return layouts.NewCategoriesLayout(m,
		html.NewView(
			paginationComponent,
			e.Div(a.Class("items is-multiline")).Children(
				items...,
			),
			paginationComponent,
		),
	)
}
