package catalog

import (
	"github.com/divilla/golastore/internal/view_model"
	"github.com/divilla/golastore/internal/web/catalog/catalog_service"
	"github.com/divilla/golastore/internal/web/views/layouts"
	"github.com/divilla/golastore/pkg/html"
	"github.com/divilla/golastore/pkg/html/a"
	"github.com/divilla/golastore/pkg/html/d"
	"github.com/divilla/golastore/pkg/html/e"
)

func NewCategoryView(m *catalog_service.CategoryModel) html.IView {
	var links, columns []html.Renderer
	p := m.Pagination
	if p.ShowSideNavigation() {
		links = append(links, makeLink(p.FirstLink(), e.Span(a.Class("mdi mdi-page-first")), p))
	}
	if p.ShowSideNavigation() {
		links = append(links, makeLink(p.PreviousLink(), e.Span(a.Class("mdi mdi-chevron-left")), p))
		links = append(links, e.Li().Children(e.Span(a.Class("pagination-ellipsis")).Text("&hellip;")))
	}
	for _, link := range p.NumberedLinks() {
		links = append(links, makeLink(link, nil, p))
	}
	if p.ShowSideNavigation() {
		links = append(links, e.Li().Children(e.Span(a.Class("pagination-ellipsis")).Text("&hellip;")))
		links = append(links, makeLink(p.NextLink(), e.Span(a.Class("mdi mdi-chevron-right")), p))
	}
	if p.ShowSideNavigation() {
		links = append(links, makeLink(p.LastLink(), e.Span(a.Class("mdi mdi-page-last")), p))
	}

	for _, item := range m.Products {
		columns = append(columns,
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
			),
		)
	}

	return layouts.NewCategoriesLayout(m, html.NewView(
		e.Nav(a.Class("pagination"), a.Role("navigation"), a.AriaLabel("pagination")).Children(
			e.Ul(a.Class("pagination-list")).Children(
				links...,
			),
		),
		e.Div(a.Class("columns is-multiline")).Children(
			columns...,
		),
	))
}

func makeLink(page view_model.PageLinkVM, text html.Renderer, p *view_model.PaginationVM) html.Renderer {
	if text == nil {
		text = d.Block(page.StrPage)
	}

	if page.Current {
		return e.Li().Children(
			e.A(a.Class("pagination-link is-current"), a.AriaLabel("Page "+page.StrPage), a.AriaCurrent("page")).Children(
				text,
			),
		)
	}

	if page.Disabled {
		return e.Li().Children(
			e.A(a.Class("pagination-link is-disabled"), a.AriaLabel("Goto page "+page.StrPage)).Children(
				text,
			),
		)
	}

	return e.Li().Children(
		e.A(a.Href(p.Url(page.StrPage)), a.Class("pagination-link"), a.AriaLabel("Goto page "+page.StrPage)).Children(
			text,
		),
	)
}
