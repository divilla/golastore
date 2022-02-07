package catalog

import (
	"fmt"
	"github.com/divilla/golastore/framework/pagination"
	"github.com/divilla/golastore/internal/web/catalog/catalog_service"
	"github.com/divilla/golastore/internal/web/views/layouts"
	"github.com/divilla/golastore/pkg/html"
	"github.com/divilla/golastore/pkg/html/a"
	"github.com/divilla/golastore/pkg/html/d"
	"github.com/divilla/golastore/pkg/html/e"
)

func NewCategoryView(model *catalog_service.CategoryModel) html.IView {
	var links, columns []html.Renderer
	pages := pagination.New(model.CurrentPage(), model.TotalPages(), 4)
	if pages.ShowFirstLast() {
		page, enabled := pages.First()
		links = append(links, pageLink(model.SelectedSlug(), e.Span(a.Class("material-icons")).Text("first_page"), page, enabled, false))
	}
	if pages.ShowPreviousNext() {
		page, enabled := pages.Previous()
		links = append(links, pageLink(model.SelectedSlug(), e.Span(a.Class("material-icons")).Text("chevron_left"), page, enabled, false))
		links = append(links, e.Li().Children(e.Span(a.Class("pagination-ellipsis")).Text("&hellip;")))
	}
	for _, page := range pages.Links() {
		links = append(links, pageLink(model.SelectedSlug(), nil, page, true, pages.IsCurrent(page)))
	}
	if pages.ShowPreviousNext() {
		page, enabled := pages.Next()
		links = append(links, e.Li().Children(e.Span(a.Class("pagination-ellipsis")).Text("&hellip;")))
		links = append(links, pageLink(model.SelectedSlug(), e.Span(a.Class("material-icons")).Text("chevron_right"), page, enabled, false))
	}
	if pages.ShowFirstLast() {
		page, enabled := pages.Last()
		links = append(links, pageLink(model.SelectedSlug(), e.Span(a.Class("material-icons")).Text("last_page"), page, enabled, false))
	}

	for _, item := range model.ProductsList() {
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

	return layouts.NewCategoriesLayout(model, html.NewView(
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

func pageLink(slug string, text html.Renderer, page int64, enabled, current bool) html.Renderer {
	var label string
	href := fmt.Sprintf("/c/%s/%d", slug, page)
	if text == nil {
		text = d.Block(fmt.Sprintf("%d", page))
	}

	if current {
		label = fmt.Sprintf("Page %d", page)
		return e.Li().Children(
			e.A(a.Class("pagination-link is-current"), a.AriaLabel(label), a.AriaCurrent("page")).Children(
				text,
			),
		)
	}

	label = fmt.Sprintf("Goto page %d", page)
	if enabled {
		return e.Li().Children(
			e.A(a.Href(href), a.Class("pagination-link"), a.AriaLabel(label)).Children(
				text,
			),
		)
	}

	return e.Li().Children(
		e.A(a.Class("pagination-link"), a.AriaLabel(label)).Children(
			text,
		),
	)
}
