package pagination

import (
	"github.com/divilla/golastore/internal/view_model"
	"github.com/divilla/golastore/pkg/html"
	"github.com/divilla/golastore/pkg/html/a"
	"github.com/divilla/golastore/pkg/html/d"
	"github.com/divilla/golastore/pkg/html/e"
	"strconv"
)

type (
	IData interface {
		Pagination() *view_model.Pagination
	}

	component struct {
		view_model.Pagination
	}

	pageLink struct {
		number   string
		url      string
		current  bool
		disabled bool
	}
)

func New(data IData) *html.Component {
	c := &component{
		*data.Pagination(),
	}

	return c.html()
}

func (c *component) html() *html.Component {
	var links []html.Renderer

	if c.showSideNavigation() {
		links = append(links, makeElm(c.firstLink(), e.Span(a.Class("mdi mdi-page-first"))))
	}
	links = append(links, makeElm(c.previousLink(), e.Span(a.Class("mdi mdi-chevron-left"))))
	links = append(links, e.Li().Children(e.Span(a.Class("pagination-ellipsis")).Text("&hellip;")))
	for _, link := range c.numberedLinks() {
		links = append(links, makeElm(link, nil))
	}
	links = append(links, e.Li().Children(e.Span(a.Class("pagination-ellipsis")).Text("&hellip;")))
	links = append(links, makeElm(c.nextLink(), e.Span(a.Class("mdi mdi-chevron-right"))))
	if c.showSideNavigation() {
		links = append(links, makeElm(c.lastLink(), e.Span(a.Class("mdi mdi-page-last"))))
	}

	return html.NewComponent(
		e.Div(a.Class("columns is-multiline")).Children(
			e.Div(a.Class("column is-full-mobile is-8-widescreen")).Children(
				e.Nav(a.Class("pagination"), a.Role("navigation"), a.AriaLabel("pagination")).Children(
					e.Ul(a.Class("pagination-list")).Children(
						links...,
					),
				),
			),
			e.Div(a.Class("column is-half-mobile is-2-widescreen")).Children(
				e.Div(a.Class("select")).Children(
					e.Select().Children(
						e.Option(a.Value("default")).Text("Default Sort"),
						e.Option(a.Value("price-desc")).Text("Price Descending"),
						e.Option(a.Value("price-asc")).Text("Price Ascending"),
						e.Option(a.Value("name-desc")).Text("Name Descending"),
						e.Option(a.Value("name-asc")).Text("Name Ascending"),
					),
				),
			),
			e.Div(a.Class("column is-half-mobile is-2-widescreen")).Children(
				e.Div(a.Class("select")).Children(
					e.Select().Children(
						e.Option(a.Value("default")).Text("Default Sort"),
						e.Option(a.Value("price-desc")).Text("Price Descending"),
						e.Option(a.Value("price-asc")).Text("Price Ascending"),
						e.Option(a.Value("name-desc")).Text("Name Descending"),
						e.Option(a.Value("name-asc")).Text("Name Ascending"),
					),
				),
			),
		),
	)
}

func (c *component) url(page int64) string {
	return c.URLBuilder(page)
}

func (c *component) showSideNavigation() bool {
	return c.TotalPages > c.TotalLinks
}

func (c *component) link(page int64) *pageLink {
	number := strconv.FormatInt(page, 10)
	return &pageLink{
		number:  number,
		url:     c.url(page),
		current: page == c.CurrentPage,
	}
}

func (c *component) firstLink() *pageLink {
	link := c.link(1)
	if c.CurrentPage == 1 {
		link.disabled = true
		link.current = false
	}

	return link
}

func (c *component) previousLink() *pageLink {
	link := c.link(c.CurrentPage - 1)
	if c.CurrentPage == 1 {
		link.disabled = true
	}

	return link
}

func (c *component) numberedLinks() []*pageLink {
	from := int64(1)
	if c.CurrentPage >= c.TotalPages-c.LinksSpread {
		from = c.TotalPages - c.TotalLinks + 1
	} else if c.CurrentPage > c.LinksSpread+1 {
		from = c.CurrentPage - c.LinksSpread
	}

	links := make([]*pageLink, c.TotalLinks)
	for i := int64(0); i < c.TotalLinks; i++ {
		links[i] = c.link(from + i)
	}

	return links
}

func (c *component) nextLink() *pageLink {
	link := c.link(c.CurrentPage + 1)
	if c.CurrentPage == c.TotalPages {
		link.disabled = true
	}

	return link
}

func (c *component) lastLink() *pageLink {
	link := c.link(c.TotalPages)
	if c.CurrentPage == c.TotalPages {
		link.disabled = true
		link.current = false
	}

	return link
}

func makeElm(page *pageLink, text html.Renderer) html.Renderer {
	if text == nil {
		text = d.Block(page.number)
	}

	if page.current {
		return e.Li().Children(
			e.A(a.Class("pagination-link is-current"), a.AriaLabel("Page "+page.number), a.AriaCurrent("page")).Children(
				text,
			),
		)
	}

	if page.disabled {
		return e.Li().Children(
			e.A(a.Class("pagination-link is-disabled"), a.AriaLabel("Goto page "+page.number)).Children(
				text,
			),
		)
	}

	return e.Li().Children(
		e.A(a.Href(page.url), a.Class("pagination-link"), a.AriaLabel("Goto page "+page.number)).Children(
			text,
		),
	)
}
